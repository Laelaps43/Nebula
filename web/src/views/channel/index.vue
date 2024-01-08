<template>
  <div class="node-conatiner">
    <h2 class="font18 marT13 rowSC link" @click="handleGoBack">
      <LeftOutlined />
      <span class="marL10">{{ '设备列表' }}</span>
    </h2>
    <a-divider />
    <h2 class="nc_title font18">设备通道</h2>
    <Table
      ref="ELRef"
      :url="fetchApi.channel_page"
      :columns="columns"
      :hiddenFilter="false"
      :button="tableFilterButton"
      :actions="tableActions"
      :status-list="ChannelStatusList"
      :page-param="pageParam"
    />
    <!-- 更新通道 -->
    <Modal
      v-bind="modalUpdateChannel"
      @cancel="handleUpdateChannelCancel()"
      @ok="handelUpdateChannelSubmit()"
    >
      <a-form
        ref="FormUpdateRef"
        :model="formUpdateChannel"
        :rules="rules"
        :label-col="labelCol"
        :wrapper-col="wrapperCol"
      >
        <a-form-item label="名称:" name="name">
          <a-input v-model:value="formUpdateChannel.name" placeholder="请输入设备名称" />
        </a-form-item>
      </a-form>
    </Modal>

    <!-- 新建通道 -->
    <Modal
      v-bind="modalCreateChannel"
      @cancel="handleCreateChannelCancel"
      @ok="handleCreateChannelSubmit"
    >
      <a-form
        ref="FormRef"
        :model="formCreateChannel"
        :rules="rules"
        :label-col="labelCol"
        :wrapper-col="wrapperCol"
      >
        <a-form-item label="名称:" name="name">
          <a-input v-model:value="formCreateChannel.name" placeholder="请输入设备名称" />
        </a-form-item>
        <a-form-item label="设备编号:" name="deviceId">
          <div style="display: flex; align-items: center"
            ><span>{{ formCreateChannel.channelId }}</span>
            <sync-outlined
              @click="refreshDeviceId"
              two-tone-color="#eb2f96"
              style="margin-left: 8px"
            />
          </div>
        </a-form-item>
      </a-form>
    </Modal>
    <player
      ref="playerRef"
      :options="videoOptions"
      :visible="playVisible"
      @cancel="handlePlayCancel"
    />
  </div>
</template>

<script setup lang="ts">
  import fetchApi from '/@/api/channel';
  import { AuthEnum } from '/@/enums/authEnum';
  import { ref } from 'vue';
  import { LeftOutlined, SyncOutlined } from '@ant-design/icons-vue';
  import { ChannelStatusList, columns } from '/@/views/channel/constant';
  import { Ref } from 'vue/dist/vue';
  import { Rule } from 'ant-design-vue/es/form';
  import { useMessage } from '/@/hooks/useMessage';
  import { FormInstance } from 'ant-design-vue';
  import { VideoRequestPayload } from '/@/api/channel/model';

  const route = useRoute();
  const router = useRouter();

  const pageParam = {
    deviceId: route.params.id,
  };

  const playerRef = ref(null);

  const labelCol = { style: { width: '110px' } };
  const wrapperCol = { span: 17 };

  const FormRef = ref<FormInstance>();
  const FormUpdateRef = ref<FormInstance>();
  const ELRef = ref<{ refresh: () => void }>();

  const { createMessage } = useMessage();

  const refresh = () => ELRef.value?.refresh();

  const handleGoBack = () => {
    router.back();
  };

  const tableFilterButton = reactive({
    type: 'primary',
    label: '新增通道',
    auth: AuthEnum.channel_create,
    onClick: () => {
      modalCreateChannel.title = '新增设备';
      modalCreateChannel.okText = '创建';
      modalCreateChannel.visible = true;
      getGenerateChannelId();
    },
  });

  // modal
  const modalCreateChannel = reactive({
    loading: false,
    visible: false,
    title: '创建通道',
    okText: '创建',
  });
  const modalUpdateChannel = reactive({
    loading: false,
    visible: false,
    title: '更新通道',
    okText: '更新',
  });
  const formUpdateChannel: Ref<{
    channelId: string;
    name: string;
  }> = ref({
    channelId: '',
    name: '',
  });
  const handleUpdateChannelCancel = () => {
    modalUpdateChannel.visible = false;
    FormUpdateRef.value?.resetFields();
  };
  const handelUpdateChannelSubmit = () => {
    FormUpdateRef.value?.validate().then(async () => {
      modalUpdateChannel.loading = true;
      const res = await fetchApi.channel_update({
        ...formUpdateChannel.value,
      });
      modalUpdateChannel.loading = false;
      modalUpdateChannel.visible = false;
      if (res) {
        createMessage.success('更新通道成功');
        handleUpdateChannelCancel();
        refresh();
      } else {
        createMessage.error('更新通道失败');
      }
    });
  };
  const formCreateChannel: Ref<{
    channelId: string;
    name: string;
  }> = ref({
    channelId: '',
    name: '',
  });
  // form
  const rules: Record<string, Rule[]> = {
    name: [{ required: true, trigger: 'blur', message: '请输入通道名称' }],
  };

  const handleCancel = () => {
    modalCreateChannel.visible = false;
    FormRef.value?.resetFields();
  };

  const handlePlayCancel = () => {
    playVisible.value = false;
  };

  const tableActions = reactive([
    {
      label: '播放',
      auth: AuthEnum.channel_show,
      onClick: async (row) => {
        await handlePlay({
          channelId: row.channelId,
          deviceId: String(route.params.id),
        });
      },
    },
    {
      label: '录制',
      enable: true,
      popConfirm: {
        title: '确认录制吗？',
        onConfirm: async (row) => {
          const result = await fetchApi.video_record({
            channelId: row.channelId,
            deviceId: String(route.params.id),
          });
          if (result) {
            createMessage.success('录制成功，请转到远端录像');
            refresh();
          }
        },
      },
      auth: AuthEnum.device_delete,
    },
    {
      label: '编辑',
      auth: AuthEnum.device_update,
      onClick: async (row) => {
        formUpdateChannel.value.channelId = row.channelId;
        modalUpdateChannel.visible = true;
      },
    },
    {
      label: '删除',
      enable: true,
      popConfirm: {
        title: '确认删除吗？',
        onConfirm: async (row) => {
          console.log('row', row);
          const result = await fetchApi.channel_delete(row.channelId);
          if (result) {
            createMessage.success('删除成功');
            refresh();
          }
        },
      },
      auth: AuthEnum.device_delete,
    },
  ]);

  const getGenerateChannelId = async () => {
    const res = await fetchApi.channel_generate();
    if (res) {
      formCreateChannel.value.channelId = res;
    } else {
      createMessage.error('生成通道Id失败');
    }
  };
  const refreshDeviceId = () => {
    getGenerateChannelId();
  };
  const handleCreateChannelCancel = () => {
    modalCreateChannel.visible = false;
    FormRef.value?.resetFields();
  };
  const handleCreateChannelSubmit = () => {
    FormRef.value?.validate().then(async () => {
      modalCreateChannel.loading = true;
      const res = await fetchApi.channel_create({
        name: formCreateChannel.value.name,
        channelId: formCreateChannel.value.channelId,
        deviceId: route.params.id[0],
      });
      modalCreateChannel.loading = false;
      modalCreateChannel.visible = false;
      if (res) {
        createMessage.success('创建通道成功');
        handleCancel();
        refresh();
      } else {
        createMessage.error('创建通道失败');
      }
    });
  };

  const videoOptions = ref({
    autoplay: true,
    controls: true,
    src: '',
    type: 'application/x-mpegURL',
  });
  const playVisible = ref(false);
  const handlePlay = async (payload: VideoRequestPayload) => {
    const res = await fetchApi.video_play(payload);
    if (res) {
      videoOptions.value.src = res.HTTP;
      playVisible.value = true;
    } else {
      createMessage.error('点播失败');
    }
  };
</script>

<style scoped lang="less"></style>
