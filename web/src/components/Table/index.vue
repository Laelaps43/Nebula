<template>
  <div class="table-component">
    <!-- filter -->
    <TableFilter
      :button="tableFilterButton"
      :items="tableFilterItems"
      :model="tableFilterModel"
      :hiddenFilter="hiddenFilter"
      :pagination="null"
      @onSearch="onSearch"
    />
    <!-- table -->
    <a-table
      :class="['ant-table-striped', { border: hasBordered }]"
      :rowClassName="(_, index) => (index % 2 === 1 ? 'table-striped' : '')"
      :dataSource="dataSource"
      :columns="columns"
      :rowKey="(record) => record.key || record.id"
      :pagination="pagination"
      :loading="loading"
      @change="handleTableChange"
      :scroll="scroll"
    >
      <template #bodyCell="{ column, text, record }">
        <template v-if="column.key === 'status'">
          <a-tag :bordered="false" :color="getStatusColor(record.status || record.enable || record.isRecording)"
            >{{ getStatusLabel(record.status || record.enable || record.isRecording) }}
          </a-tag>
        </template>
        <template v-if="column.key === 'role'">
          <cascader :role-list="roleList" :user-role="getUserRole(record)" />
        </template>
        <template v-if="column.key === 'toDate'">
          <span>{{ text ? formatDate(text) : '-' }}</span>
        </template>
        <template v-if="column.key === 'toDateTime'">
          <span>{{ text ? formatDate(text, 'time') : '-' }}</span>
        </template>
        <!-- 函数式写法自定义 操作列 -->
        <template v-if="column.key === 'action'">
          <template v-for="(action, index) in getActions" :key="`${index}-${action.label}`">
            <!-- 气泡确认框 -->
            <a-popconfirm
              v-if="action.enable && action.permission"
              :title="action?.title"
              @confirm="action?.onConfirm(record)"
              @cancel="action?.onCancel(record)"
            >
              <a @click.prevent="() => {}" :type="action.type">{{ action.label }}</a>
            </a-popconfirm>
            <span v-else-if="action.permission">
              <!-- 按钮 -->
              <a @click="action?.onClick(record)" :type="action.type">{{ action.label }}</a>
            </span>
            <!-- 分割线 -->
            <a-divider type="vertical" v-if="index < getActions.length - 1" />
          </template>
        </template>
      </template>
    </a-table>
  </div>
</template>
<script lang="ts">
  import dayjs from 'dayjs';
  import { formatToDate, formatToDateTime } from '/@/utils/dateUtil';
  import { usePermission } from '/@/hooks/usePermission';
  import { TablePaginationConfig } from 'ant-design-vue';
  import type { CascaderProps } from 'ant-design-vue';
  import { FilterValue } from 'ant-design-vue/es/table/interface';

  export default defineComponent({
    props: [
      'bordered',
      'hiddenFilter',
      'url' /* 请求接口 promise */,
      'columns' /* Table组件：columns 不包含操作列 */,
      'actions' /* Table组件：操作列 */,
      'button' /* Filter筛选列组件：交互按钮 */,
      'items' /* Filter筛选列组件：包含的项 */,
      'model' /* Filter筛选列组件：form model */,
      'resKey',
      'scroll',
      'statusList',
      'pageParam',
      'roleList',
    ],

    setup(props) {
      const { hasPermission } = usePermission();

      const options: CascaderProps['options'] = [
        {
          label: 'Light',
          value: 'light',
          children: new Array(20)
            .fill(null)
            .map((_, index) => ({ label: `Number ${index}`, value: index })),
        },
        {
          label: 'Bamboo',
          value: 'bamboo',
          children: [
            {
              label: 'Little',
              value: 'little',
              children: [
                {
                  label: 'Toy Fish',
                  value: 'fish',
                },
                {
                  label: 'Toy Cards',
                  value: 'cards',
                },
                {
                  label: 'Toy Bird',
                  value: 'bird',
                },
              ],
            },
          ],
        },
      ];

      const valueT = ref<string[]>([]);
      const dataSource = ref([]);
      const loading = ref(false);
      const current = ref(1);
      const pageSize = ref(10);
      const total = ref(0);
      const hasBordered = computed(() => props.bordered ?? true);
      const roleList = props.roleList;

      const pagination = computed(() => ({
        total: total.value,
        current: current.value,
        pageSize: pageSize.value,
        showQuickJumper: true,
        showSizeChanger: true,
        showTotal: () => h('span', {}, `共 ${total.value} 条`),
      }));

      onMounted(() => {
        initTableData();
      });

      const handleTableChange = (
        pag: TablePaginationConfig,
        filters: Record<string, FilterValue | null>,
        sorter: any,
      ) => {
        pageSize.value = pag.pageSize;
      };
      const initTableData = () => {
        props
          .url({
            limit: pageSize.value,
            page: current.value,
            ...props.pageParam,
          })
          .then((res: any) => {
            dataSource.value = res.list;
            total.value = res.total;
          });
      };
      const refresh = () => {
        initTableData();
      };
      // action 操作列
      const getActions = computed(() => {
        return (
          (toRaw(props.actions) || [])
            // .filter((action) => hasPermission(action.auth))
            .map((action) => {
              const { popConfirm } = action;
              return {
                type: 'link',
                ...action,
                ...(popConfirm || {}),
                enable: action.enable,
                permission: hasPermission(action.auth),
              };
            })
        );
      });

      // filter
      const tableFilterModel = computed(() => props.model);
      const tableFilterButton = computed(() => props.button);
      const tableFilterItems = computed(() => props.items);
      // const hiddenFilter = computed(() => props.hiddenFilter);
      const onSearch = () => {
        const args = toRaw(tableFilterModel.value) || {};

        // 日期格式处理
        if (args) {
          Object.keys(args).map((key) => {
            if (args[key] && dayjs.isDayjs(args[key])) {
              args[key] = formatToDate(args[key]);
            }
          });
        }
      };

      // 日期格式化
      const formatDate = (val: string, type: 'date' | 'time' = 'date') => {
        const formatFn = type === 'date' ? formatToDate : formatToDateTime;
        return val.length === 10 ? formatFn(Number(val) * 1000) : formatFn(val);
      };

      const getStatusLabel = computed(() => {
        return (status: any) => {
          return (props.statusList as []).find((item) => item?.status == status)?.label || '-';
        };
      });
      const getStatusColor = computed(() => {
        return (status: any) => {
          return (props.statusList as []).find((item) => item?.status == status)?.color || '-';
        };
      });

      const getUserRole = computed(() => (user) => {
        return user.roles?.map((role) => [role?.id]).filter(Boolean) || [];
      });
      return {
        dataSource,
        loading,
        pagination,
        hasBordered,
        handleTableChange,
        // run,
        refresh,
        valueT,
        options,
        getActions,
        tableFilterModel,
        tableFilterButton,
        tableFilterItems,
        onSearch,
        formatDate,
        getStatusLabel,
        getStatusColor,
        getUserRole,
      };
    },
  });
</script>
<style lang="less" scoped>
  .ant-table-striped :deep(.table-striped) td {
    background-color: #fafafa;
  }

  .ant-table-striped :deep(.ant-table-pagination.ant-pagination) {
    margin: 30px auto;
    width: 100%;
    text-align: center;

    .ant-pagination-prev,
    .ant-pagination-next {
      .anticon {
        vertical-align: 1.5px;
      }
    }
  }

  .ant-table-striped :deep(.ant-pagination-item-active) {
    background: #3860f4;

    a {
      color: #ffffff;
    }
  }

  .border {
    border: 0.5px solid rgba(210, 210, 210, 0.5);
  }
</style>
