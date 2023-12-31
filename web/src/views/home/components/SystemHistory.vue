<template>
  <div class="systemHistory">
    <div class="systemHistory_header">
      <h2 class="font18">服务器状态</h2>
    </div>
    <a-card
      :bordered="true"
      :loading="loading"
      :body-style="{ padding: 0 }"
      class="systemHistory_chart"
    >
      <div class="chart-box" ref="chartRef"></div>
    </a-card>
  </div>
</template>
<script setup lang="ts">
  import type { Ref } from 'vue';
  import { useECharts } from '/@/hooks/useECharts';
  import { useHomeStore } from '/@/store/modules/home';

  const store = useHomeStore();

  onMounted(() => {
    store.fetchSystemInfo();
  });

  let interval = setInterval(store.fetchSystemInfo, 60000);

  onUnmounted(() => {
    clearInterval(interval);
  });

  let data = {
    TimeList: [],
    CPUList: [],
    DiskList: [],
    MemList: [],
    DownList: [],
    UpList: [],
  };

  let dataSource = computed(() => store.getSystemInfo || data);

  const props = defineProps({
    loading: Boolean,
  });

  const chartRef = ref<HTMLDivElement | null>(null);
  const { setOptions } = useECharts(chartRef as Ref<HTMLDivElement>);

  watch(
    () => [props.loading, dataSource.value],
    () => {
      if (props.loading) {
        return;
      }

      setOptions({
        grid: {
          left: 0,
          right: 10,
          top: 30,
          bottom: 0,
          containLabel: true,
        },
        legend: {
          show: true,
          left: 'center',
          data: ['CPU', 'Disk', 'Memory', 'Download', 'Upload'],
        },
        tooltip: {
          trigger: 'axis',
          type: 'axis',
        },
        xAxis: {
          type: 'category',
          data: dataSource.value.TimeList,
          boundaryGap: false,
          axisLabel: {
            rotate: 40,
            margin: 16,
            color: '#999999',
          },
        },
        yAxis: [
          {
            type: 'value',
            position: 'left',
            axisLabel: {
              formatter: '{value} %',
            },
          },
          {
            type: 'value',
            name: 'Mbps',
            position: 'right',
            splitLine: {
              show: false,
            },
          },
        ],
        series: [
          {
            data: dataSource.value.CPUList,
            type: 'line',
            name: 'CPU',
            yAxisIndex: 0,
            smooth: true,
          },
          {
            data: dataSource.value.DownList,
            type: 'line',
            name: 'Download',
            yAxisIndex: 1,
            smooth: true,
          },
          {
            data: dataSource.value.DiskList,
            type: 'line',
            name: 'Disk',
            yAxisIndex: 0,
            smooth: true,
          },
          {
            data: dataSource.value.MemList,
            type: 'line',
            name: 'Memory',
            yAxisIndex: 0,
            smooth: true,
          },
          {
            data: dataSource.value.UpList,
            type: 'line',
            name: 'Upload',
            yAxisIndex: 1,
            smooth: true,
          },
        ],
      });
    },
    { immediate: true },
  );
</script>
<style lang="less" scoped>
  .systemHistory {
    background: #ffffff;

    &_chart {
      .chart-box {
        width: 100%;
        height: 245px;
      }
    }
  }
</style>
