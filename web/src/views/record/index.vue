<template>
  <div class="node-conatiner">
    <h2 class="nc_title font18">远端录像</h2>
    <Table
      ref="ELRef"
      :url="fetchApi.recordPage"
      :columns="columns"
      :hiddenFilter="false"
      :actions="tableActions"
      :status-list="RecordStatusList"
    />
    <calendar
      ref="calendarRef"
      :title="calendarTitle"
      :visible="calendarVisible"
      :stream="calendarStream"
      @cancel="cancelCalendar"
    />
  </div>
</template>
<script setup lang="ts">
  import { columns, RecordStatusList } from './constant';
  import fetchApi from '/@/api/record';
  import { reactive } from 'vue';
  import { AuthEnum } from '/@/enums/authEnum';

  const calendarRef = ref(null);
  const calendarTitle = ref('');
  const calendarVisible = ref(false);
  const calendarStream = ref('');

  const tableActions = reactive([
    {
      label: '查看',
      onClick: async (row) => {
        calendarTitle.value = row.channelName;
        calendarStream.value = row.stream;
        calendarVisible.value = true;
        await nextTick(() => {
          calendarRef.value?.initCalendar();
        });
      },
      auth: AuthEnum.record_range,
    },
  ]);

  const cancelCalendar = () => {
    calendarVisible.value = false;
  };
</script>
<style lang="less" scoped>
  .node-conatiner {
    .nc_title {
      margin-top: 6px;
      margin-bottom: 30px;
    }
  }
</style>
