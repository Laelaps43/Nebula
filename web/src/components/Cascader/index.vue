<!-- ChildComponent.vue -->
<template>
  <a-cascader
    v-model:value="value"
    style="width: 100%"
    multiple
    max-tag-count="responsive"
    :allowClear="false"
    :options="roleList"
    placeholder="请选择角色"
    :show-checked-strategy="Cascader.SHOW_CHILD"
  />
</template>

<script setup lang="ts">
  import { defineProps } from 'vue';
  import { Cascader } from 'ant-design-vue';

  const { userRole, roleList } = defineProps(['userRole', 'roleList']);
  const value = userRole;
</script>
