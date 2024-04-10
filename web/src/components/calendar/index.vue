<template>
  <a-modal
    :title="title"
    :width="'45vw'"
    :visible="visible"
    wrap-class-name="full-modal"
    @cancel="cancel"
    :footer="null"
  >
    <a-calendar :fullscreen="false" :mode="'month'" @change="onChange" @panelChange="onPanelChange">
      <template #dateCellRender="{ current }">
        <div v-if="getDateVideoStatus(current)" style="display: flex; justify-content: center">
          <div style="width: 5px; height: 5px; background: #ffc533; border-radius: 3px"></div>
        </div>
      </template>
    </a-calendar>
    <div style="display: flex; justify-content: right">
      <a-button v-if="hasPermission(AuthEnum.record_details)" @click="toCheckVideo()">查看</a-button>
    </div>
  </a-modal>
  <player
    :ref="(el) => (playerRef = el)"
    :visible="playerVisible"
    :options="videoOptions"
    @cancel="playerCancel"
    @changePlay="changePlay"
    :is-live="false"
    :select-options="selectOptions"
  />
</template>
<script lang="ts">
  import dayjs, { Dayjs } from 'dayjs';
  import fetchApi from '/@/api/record';
  import { ref } from 'vue';
  import { useMessage } from '/@/hooks/useMessage';
  import { usePermission } from "/@/hooks/usePermission";
  import { AuthEnum } from "/@/enums/authEnum";

  export default defineComponent({
    computed: {
      AuthEnum() {
        return AuthEnum
      }
    },
    props: {
      title: {
        type: String,
        default: '',
      },
      visible: {
        type: Boolean,
        default: false,
      },
      loading: {
        type: Boolean,
        default: false,
      },
      fetchFunc: {
        type: Function,
        default: fetchApi.recordRange,
      },
      stream: {
        type: String,
        default: '',
      },
    },
    emits: ['cancel'],
    setup(props, { emit }) {
      const value = ref<string>();
      const isRecorded = ref();

      const { hasPermission } = usePermission();

      const initCalendar = () => {
        onPanelChange(dayjs(), 'month');
      };

      const getDateVideoStatus = (current: Dayjs) => {
        if (!isRecorded.value) {
          return false;
        }
        return isRecorded.value?.find(
          (date) => dayjs(date).format('YYYY-MM-DD') === current.format('YYYY-MM-DD'),
        );
      };

      const onPanelChange = (data: Dayjs, mode: string) => {
        if (mode === 'year') {
          return;
        }
        let beginData = data.subtract(42, 'day').format('YYYY-MM-DD');
        let endData = data.add(42, 'day').format('YYYY-MM-DD');
        getVideoRange(beginData, endData);
      };

      const handleCancel = () => emit('cancel');

      const cancel = () => {
        emit('cancel');
      };

      const { createMessage } = useMessage();
      const playerVisible = ref(false);
      const playerRef = ref(null);
      const toCheckVideo = async () => {
        let res = await fetchApi.recordSelect(
          value.value || dayjs().format('YYYY-MM-DD'),
          props.stream,
        );
        if (res) {
          selectOptions.value = Array.isArray(res) ? res : [];
          if (selectOptions.value.length > 0) {
            playerRef.value.setValue(selectOptions.value[0].value);
            playVideoRecord(selectOptions.value[0].value);
            playerVisible.value = true;
          } else {
            createMessage.error('暂无录像');
          }
        } else {
          createMessage.error('查看失败');
        }
      };
      const changePlay = (id: number) => {
        playVideoRecord(id);
      };
      const playVideoRecord = (id: number) => {
        fetchApi.recordPlay(id, props.stream).then((res) => {
          videoOptions.value.src = res
        });
      };

      const onChange = (data: Dayjs) => {
        value.value = data.format('YYYY-MM-DD');
      };

      const getVideoRange = (start: string, end: string) => {
        props
          .fetchFunc({
            start: start,
            end: end,
            stream: props.stream,
          })
          .then((res) => {
            isRecorded.value = Array.isArray(res) ? res : [];
          });
      };

      const videoOptions = ref({
        autoplay: true,
        controls: true,
        src: '',
        type: 'application/x-mpegURL',
      });

      const selectOptions = ref();
      const playerCancel = () => {
        playerVisible.value = false;
      };
      return {
        value,
        cancel,
        onChange,
        playerRef,
        changePlay,
        handleCancel,
        videoOptions,
        toCheckVideo,
        initCalendar,
        playerCancel,
        selectOptions,
        playerVisible,
        onPanelChange,
        hasPermission,
        getDateVideoStatus,
      };
    },
  });
</script>

<style lang="less">
  .full-modal {
    .ant-modal {
      max-width: 80%;
    }

    .ant-modal-content {
      display: flex;
      flex-direction: column;
    }

    .ant-modal-body {
      flex: 1;
    }
  }
</style>
