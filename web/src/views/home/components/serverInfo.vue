<template>
  <div class="dataOverview">
    <div class="font18">设备数据总览</div>
    <section class="dataOverview_row">
      <div
        v-for="item in data"
        :key="item.label"
        :span="12"
        :style="{ background: item.bgColor }"
        class="dataOverview_col"
      >
        <div><img :src="item.img" alt="" /></div>
        <div>
          <h2>{{ item.label }}</h2>
          <div>
            <p>媒体服务地址:192.168.2.214</p>
            <p>RTP端口:8081</p>
            <p>restful端口:80</p>
            <p>TCP会话:8</p>
            <p>UDP会话:10</p>
            <p>最近心跳时间:2024-01-09 22:20:11</p>
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

  const data = ref([
    {
      img: img1,
      label: '媒体服务器',
      value: 0,
      key: 'onlineDevice',
      bgColor: 'linear-gradient(90deg, #4394FF 0%, #60BBFD 100%)',
    },
    {
      img: img2,
      label: '信令服务器',
      value: 0,
      key: 'offlineDevice',
      bgColor: 'linear-gradient(90deg, #49C4D8 0%, #6DD2A2 100%)',
    },
  ]);

  watchEffect(() => {
    if (store.overView) {
      data.value = unref(data.value).map((n) => {
        let value = store.overView?.[n.key] || 0;
        return {
          ...n,
          value,
        };
      });
    }
  });
</script>
<style lang="less" scoped>
  .dataOverview {
    &_row {
      margin-top: 24px;
      display: grid;
      grid-template-columns: 1fr 1fr;
      gap: 21px;
    }

    &_col {
      display: flex;
      border-radius: 6px;
      height: 200px;
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
