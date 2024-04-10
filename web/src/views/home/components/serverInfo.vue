<template>
  <div class="dataOverview">
    <div class="font18">设备数据总览</div>
    <section class="dataOverview_row">
      <div
        style="background: linear-gradient(90deg, #4394ff 0%, #60bbfd 100%)"
        class="dataOverview_col"
      >
        <div><img :src="img1" alt="" /></div>
        <div style="width: 100%">
          <h2>媒体服务器</h2>
          <div>
            <a-row>
              <p>媒体服务地址:{{ dataSource.mediaServerDetails.mediaServiceAddress }}</p>
            </a-row>
            <a-row>
              <p>媒体唯一标识:{{ dataSource.mediaServerDetails.mediaUniqueID }}</p>
            </a-row>
            <a-row>
              <a-col :span="12">
                <p>RTP端口:{{ dataSource.mediaServerDetails.rtpPort }}</p>
              </a-col>
              <a-col :span="12">
                <p>restful端口:{{ dataSource.mediaServerDetails.restfulPort }}</p>
              </a-col>
            </a-row>
            <a-row>
              <a-col :span="12">
                <p>RTSP端口:{{ dataSource.mediaServerDetails.rtspPort }}</p>
              </a-col>
              <a-col :span="12">
                <p>RTMP端口:{{ dataSource.mediaServerDetails.rtmpPort }}</p>
              </a-col>
            </a-row>
            <a-row>
              <a-col :span="12">
                <p>TCP会话:{{ dataSource.mediaServerDetails.tcpSessions }}</p>
              </a-col>
              <a-col :span="12">
                <p>UDP会话:{{ dataSource.mediaServerDetails.udpSessions }}</p>
              </a-col>
            </a-row>
            <a-row>
              <div style="display: flex; flex-wrap: wrap">
                <p>最近心跳时间:</p>
                <p style="margin-left: 10px">{{
                  dataSource.mediaServerDetails.lastHeartbeatTime
                }}</p>
              </div>
            </a-row>
          </div>
        </div>
      </div>
      <div
        style="background: linear-gradient(90deg, #49c4d8 0%, #6dd2a2 100%)"
        class="dataOverview_col"
      >
        <div><img :src="img2" alt="" /></div>
        <div style="width: 100%">
          <h2>服务器信息</h2>
          <div>
            <a-row>
              <p>媒体服务地址:{{ dataSource.serverDetails.serviceAddress }}</p>
            </a-row>
            <a-row>
              <div style="display: flex; flex-wrap: wrap">
                <p>SIP服务器ID:</p>
                <p style="margin-left: 10px">{{ dataSource.serverDetails.sipServerID }}</p>
              </div>
            </a-row>
            <a-row>
              <p>SIP服务器域:{{ dataSource.serverDetails.sipServerDomain }}</p>
            </a-row>
            <a-row>
              <p
                >SIP密码:<span>{{ dataSource.serverDetails.sipPassword }}</span></p
              >
            </a-row>
            <a-row>
              <div style="display: flex; flex-wrap: wrap">
                <p>系统运行时间:</p>
                <p style="margin-left: 10px">{{ getFormatTimeDifference(parseInt(dataSource.serverDetails.uptime, 10)) }}</p>
              </div>
            </a-row>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>
<script setup lang="ts">
  import img1 from '/@/assets/images/Icon_Block.png';
  import img2 from '/@/assets/images/Icon_trading.png';

  import { useHomeStore } from '/@/store/modules/home';

  const store = useHomeStore();

  onMounted(() => {
    store.fetchServerInfo();
  });
  let interval = setInterval(store.fetchSystemInfo, 60000 * 10);

  onUnmounted(() => {
    clearInterval(interval);
  });

  let dataSource = computed(() => store.getServerInfo || getDefaultServerInfo());

  let getFormatTimeDifference = computed(() => {
    return (timeDifferenceInMilliseconds: number) => {
      const totalSeconds = Math.abs(Math.floor(timeDifferenceInMilliseconds / 1000));

      const days = Math.floor(totalSeconds / (3600 * 24));
      const hours = Math.floor((totalSeconds % (3600 * 24)) / 3600);
      const minutes = Math.floor((totalSeconds % 3600) / 60);

      return `${days}天 ${hours}小时 ${minutes}分钟`;
    };
  });

  const getDefaultServerInfo = () => {
    return {
      serverDetails: {
        serviceAddress: '',
        sipServerID: '',
        sipServerDomain: '',
        sipPassword: '',
        uptime: 0,
      },
      mediaServerDetails: {
        mediaServiceAddress: '',
        mediaUniqueID: '',
        rtpPort: 0,
        restfulPort: 0,
        rtspPort: 0,
        rtmpPort: 0,
        tcpSessions: 0,
        udpSessions: 0,
        lastHeartbeatTime: '',
      },
    };
  };
</script>
<style lang="less" scoped>
  .dataOverview {
    &_row {
      margin-top: 20px;
      display: grid;
      grid-template-columns: 1fr 1fr;
      gap: 21px;
    }

    &_col {
      display: flex;
      border-radius: 6px;
      height: 205px;
      padding: 15px;

      h2 {
        font-size: 20px;
        font-family: Roboto;
        font-weight: bold;
        color: #ffffff;
        margin-bottom: 4px;
      }

      p {
        margin: 0;
        font-size: 14px;
        font-family: Source Han Sans CN;
        font-weight: 300;
        color: #ffffff;
      }
    }
  }
</style>
