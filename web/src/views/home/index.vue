<template>
  <div>
    <a-row :gutter="24">
      <a-col :span="12">
        <a-card :bordered="false">
          <DataOverview />
        </a-card>
      </a-col>
      <a-col :span="12">
        <a-card :bordered="false">
          <!-- <TradingHistory :loading="loading" /> -->
        </a-card>
      </a-col>
    </a-row>
    <a-row :gutter="24">
      <a-col :span="24">
        <a-card :bordered="false">
          <SystemHistory />
        </a-card>
      </a-col>
    </a-row>
    <!-- <Table :url="fetchApi.list" :columns="columns" :hiddenFilter="true" :scroll="{ x: 1200 }" /> -->
  </div>
</template>
<script setup lang="ts">
  import DataOverview from './components/DataOverview.vue';
  // import TradingHistory from './components/TradingHistory.vue';
  // import { columns } from './constant';
  // import fetchApi from '/@/api/home';
  import { useHomeStore } from '/@/store/modules/home';
  import SystemHistory from '/@/views/home/components/SystemHistory.vue';

  const store = useHomeStore();
  const loading = ref(false);

  onMounted(async () => {
    loading.value = true;
    await store.fetchInfo();
    loading.value = false;
  });

  onUnmounted(() => {
    store.resetState();
  });
</script>
