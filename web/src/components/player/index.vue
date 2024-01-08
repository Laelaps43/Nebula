<template>
  <a-modal
    :width="'72vw'"
    :visible="visible"
    :title="'Creamer 01'"
    @cancel="handleCancel"
    :footer="null"
    wrap-class-name="full-modal"
  >
    <div>
      <div v-if="!isLive">
        <a-select
          style="width: 12vw"
          @change="changeSelect"
          :value="value"
          :options="selectOptions"
        />
      </div>
      <div style="width: 100%; display: flex; justify-content: center">
        <video-player
          :ref="(el) => (playerRef = el)"
          :src="options.src"
          :preload="'auto'"
          :autoplay="true"
          :width="1100"
          controls
          @mounted="videoPause"
        />
      </div>
    </div>
  </a-modal>
</template>

<script lang="ts">
  import { VideoPlayer } from '@videojs-player/vue';
  import 'video.js/dist/video-js.css';

  export default {
    components: { VideoPlayer },
    props: {
      isLive: {
        type: Boolean,
        default: true,
      },
      visible: {
        type: Boolean,
        required: true,
        default: false,
      },
      options: {
        type: Object,
        required: true,
      },
      selectOptions: {
        type: Array,
      },
    },
    emits: ['cancel', 'changePlay'],
    setup(_, ctx) {
      const playerRef = ref(null);
      const handleCancel = () => {
        player.value?.pause();
        ctx.emit('cancel');
      };

      const player = ref(null);

      const value = ref();
      const videoPause = (p: any) => {
        player.value = p.player;
      };
      const setValue = (v) => {
        console.log(v);
        value.value = v;
        console.log(value.value);
      };
      const changeSelect = (v, _) => {
        value.value = v;
        ctx.emit('changePlay', v);
      };
      return {
        videoPause,
        handleCancel,
        changeSelect,
        playerRef,
        value,
        setValue,
      };
    },
  };
</script>

<style lang="less">
  .full-modal {
    .ant-modal {
      max-width: 80%;
      top: 2vh;
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
