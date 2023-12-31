<template>
  <a-modal
    :width="'72vw'"
    :visible="visible"
    :title="'Creamer 01'"
    @cancel="handleCancel"
    :footer="null"
    wrap-class-name="full-modal"
  >
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
  </a-modal>
</template>

<script lang="ts">
  import { VideoPlayer } from '@videojs-player/vue';
  import 'video.js/dist/video-js.css';

  export default {
    components: { VideoPlayer },
    props: {
      visible: {
        type: Boolean,
        required: true,
        default: false,
      },
      options: {
        type: Object,
        required: true,
        default: {},
      },
    },
    emits: ['cancel'],
    setup(props, ctx) {
      const playerRef = ref(null);
      const handleCancel = () => {
        player.value.pause();
        console.log(player.value);
        ctx.emit('cancel');
      };

      const player = ref(null);
      const videoPause = (p) => {
        player.value = p.player;
      };
      return {
        videoPause,
        handleCancel,
        playerRef,
      };
    },
  };
</script>

<style lang="less">
  .full-modal {
    .ant-modal {
      max-width: 80%;
      top: 5vh;
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
