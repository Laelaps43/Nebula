<template>
  <ConfigProvider :locale="zhCN">
    <router-view />
  </ConfigProvider>
</template>

<script lang="ts" setup>
  import { ConfigProvider } from 'ant-design-vue';
  import zhCN from 'ant-design-vue/es/locale/zh_CN';
  import { useTitle } from '/@/hooks/useTitle';

  import dayjs from 'dayjs';
  import 'dayjs/locale/zh-cn';
  dayjs.locale('zh-cn');

  useTitle();

  console.log('my config env: ', import.meta.env);
</script>
