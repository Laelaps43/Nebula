<template>
  <div class="node-conatiner">
    <h2 class="nc_title font18">设备列表</h2>
    <Table
      ref="ELRef"
      :url="fetchApi.device_page"
      :columns="columns"
      :hiddenFilter="false"
      :button="tableFilterButton"
      :actions="tableActions"
      :status-list="DeviceStatusList"
    />
    <Modal v-bind="modalState" @cancel="handleCancel" @ok="handleSubmit">
      <a-form
        ref="FormRef"
        :model="formModel"
        :rules="rules"
        :label-col="labelCol"
        :wrapper-col="wrapperCol"
      >
        <a-form-item label="名称:" name="name">
          <a-input v-model:value="formModel.name" placeholder="请输入设备名称" />
        </a-form-item>
      </a-form>
    </Modal>
    <!-- 新建设备 -->
    <Modal
      v-bind="modalCreateDevice"
      @cancel="handleCreateDeviceCancel"
      @ok="handleCreateDeviceSubmit"
    >
      <a-form
        ref="FormRef"
        :model="formCreateDevice"
        :rules="rules"
        :label-col="labelCol"
        :wrapper-col="wrapperCol"
      >
        <a-form-item label="名称:" name="name">
          <a-input v-model:value="formCreateDevice.name" placeholder="请输入设备名称" />
        </a-form-item>
        <a-form-item v-if="false" label="端口" name="port">
          <a-input-number
            v-model:value="formCreateDevice.port"
            :controls="true"
            :max="65535"
            :min="1024"
          />
        </a-form-item>
        <a-form-item label="设备编号:" name="deviceId">
          <div style="display: flex; align-items: center"
            ><span>{{ formCreateDevice.deviceId }}</span>
            <sync-outlined
              @click="refreshDeviceId"
              two-tone-color="#eb2f96"
              style="margin-left: 8px"
            />
          </div>
        </a-form-item>
      </a-form>
    </Modal>
  </div>
</template>
<script setup lang="ts">
  import { columns, DeviceStatusList } from './constant';
  import { SyncOutlined } from '@ant-design/icons-vue';
  import fetchApi from '/@/api/device';
  import { useMessage } from '/@/hooks/useMessage';
  import type { FormInstance } from 'ant-design-vue';
  import type { Rule } from 'ant-design-vue/es/form';
  import { computed, reactive, ref, Ref } from 'vue';
  import { AuthEnum } from '/@/enums/authEnum';

  const router = useRouter();
  const mockReq = (params?: any): Promise<Boolean> =>
    new Promise((resolve) => setTimeout(() => resolve(params ? !!params : true), 500));

  const { createMessage } = useMessage();
  const FormRef = ref<FormInstance>();
  const ELRef = ref<{ refresh: () => void }>();
  const refresh = () => ELRef.value?.refresh();

  const labelCol = { style: { width: '110px' } };
  const wrapperCol = { span: 17 };

  const tableFilterButton = reactive({
    type: 'primary',
    label: '新增设备',
    auth: AuthEnum.device_create,
    onClick: () => {
      getGenerateDeviceId();
      modalCreateDevice.visible = true;
    },
  });

  // 新建设备
  const modalCreateDevice = reactive({
    loading: false,
    visible: false,
    title: '新建设备',
    okText: '新建',
  });
  // modal
  const modalState = reactive({
    loading: false,
    visible: false,
    title: '修改设备信息',
    okText: '修改',
  });

  // form
  const rules: Record<string, Rule[]> = {
    name: [{ required: true, trigger: 'blur', message: '请输入设备名称' }],
    port: [{ required: true, trigger: 'blur', message: '请输入端口号' }],
  };

  const formModel = ref({
    deviceId: '-',
    name: '',
  });

  const formCreateDevice: Ref<{
    deviceId: string;
    name: string;
    port: number | null;
  }> = ref({
    deviceId: '',
    name: '',
    port: 5060,
  });

  const tableActions = reactive([
    {
      label: '通道',
      onClick: async (row) => {
        await router.push('/channel/' + row.deviceId);
      },
      auth: AuthEnum.channel_show,
    },
    {
      label: '编辑',
      auth: AuthEnum.device_update,
      onClick: async (row) => {
        modalState.visible = true;
        formModel.value.deviceId = row.deviceId;
      },
    },
    {
      label: '删除',
      enable: true,
      popConfirm: {
        title: '确认删除吗？',
        onConfirm: async (row) => {
          const res = await fetchApi.device_delete(row.deviceId);
          if (res) {
            createMessage.success('删除设备成功');
            refresh();
          }
        },
      },
      auth: AuthEnum.device_delete,
    },
  ]);

  const handleCancel = () => {
    modalState.visible = false;
    FormRef.value?.resetFields();
  };

  const handleCreateDeviceCancel = () => {
    modalCreateDevice.visible = false;
    FormRef.value?.resetFields();
  };

  const refreshDeviceId = () => {
    getGenerateDeviceId();
  };

  const getGenerateDeviceId = async () => {
    const res = await fetchApi.device_generate();
    if (res) {
      formCreateDevice.value.deviceId = res;
    } else {
      createMessage.error('生成设备Id失败');
    }
  };
  const handleCreateDeviceSubmit = () => {
    FormRef.value?.validate().then(async () => {
      modalCreateDevice.loading = true;
      const res = await fetchApi.device_create(formCreateDevice.value);
      modalCreateDevice.loading = false;
      modalCreateDevice.visible = false;
      if (res) {
        createMessage.success('创建设备成功');
        handleCancel();
        refresh();
      } else {
        createMessage.error('创建设备失败');
      }
    });
  };

  const handleSubmit = () => {
    FormRef.value
      ?.validate()
      .then(async () => {
        modalState.loading = true;
        const res = await fetchApi.device_update(formModel.value);
        modalState.loading = false;
        if (res) {
          createMessage.success('修改设备名称成功');
          handleCancel();
          refresh();
        } else {
          createMessage.error('修改设备名称失败');
        }
      })
      .catch(console.log);
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
