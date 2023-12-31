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
          <h2>{{ item.value }}</h2>
          <p>{{ item.label }}</p>
        </div>
      </div>
    </section>
  </div>
</template>
<script setup lang="ts">
  import img1 from '/@/assets/images/Icon_Block.png';
  import img2 from '/@/assets/images/Icon_trading.png';
  import img3 from '/@/assets/images/Icon_contract.png';
  import img4 from '/@/assets/images/Icon_node.png';

  import { useHomeStore } from '/@/store/modules/home';

  const store = useHomeStore();

  const data = ref([
    {
      img: img1,
      label: '在线设备',
      value: 0,
      key: 'onlineDevice',
      bgColor: 'linear-gradient(90deg, #4394FF 0%, #60BBFD 100%)',
    },
    {
      img: img2,
      label: '离线设备',
      value: 0,
      key: 'offlineDevice',
      bgColor: 'linear-gradient(90deg, #49C4D8 0%, #6DD2A2 100%)',
    },
    {
      img: img3,
      label: '通道数量',
      value: 0,
      key: 'channel',
      bgColor: 'linear-gradient(90deg, #F48D6D 0%, #F9AC67 100%)',
    },
    {
      img: img4,
      label: '正在录像',
      value: 0,
      key: 'video',
      bgColor: 'linear-gradient(90deg, #7E8BEC 0%, #C5AFF7 100%)',
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
      height: 90px;
      padding: 15px;

      h2 {
        font-size: 23px;
        font-family: Roboto;
        font-weight: bold;
        color: #ffffff;
        margin-bottom: 4px;
      }

      p {
        margin: 0;
        font-size: 12px;
        font-family: Source Han Sans CN;
        font-weight: 300;
        color: #ffffff;
      }
    }
  }
</style>
