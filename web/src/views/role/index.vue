<template>
  <div class="node-conatiner">
    <h2 class="nc_title font18">角色管理</h2>
    <div>
      <Table
        ref="ELRef"
        :url="fetchApi.ListRole"
        :columns="columns"
        :hiddenFilter="false"
        :button="tableFilterButton"
        :actions="tableActions"
        :status-list="ChannelStatusList"
      />
    </div>
    <Modal
      v-bind="modalCreateRole"
      @cancel="handelCreateRoleCancel()"
      @ok="handleCreateRoleSubmit()"
    >
      <a-form
        ref="FormCreateRoleRef"
        :model="formCreateRole"
        :rules="formCreateRoleRules"
        :label-col="labelCol"
        :wrapper-col="wrapperCol"
      >
        <a-form-item label="名称:" name="name">
          <a-input v-model:value="formCreateRole.name" placeholder="请输入角色名称" />
        </a-form-item>
        <a-form-item label="标识:" name="slug">
          <a-input v-model:value="formCreateRole.slug" placeholder="请输入角色标识" />
        </a-form-item>
        <a-form-item label="描述:" name="desc">
          <a-input v-model:value="formCreateRole.desc" placeholder="请输入角色描述" />
        </a-form-item>
        <a-form-item label="角色:" name="parentRole">
          <a-select disabled v-model:value="formCreateRole.parentId" :options="options"></a-select>
        </a-form-item>
      </a-form>
    </Modal>
    <!-- 编辑角色   -->
    <Modal
      v-bind="modalUpdateRole"
      @cancel="handelUpdateRoleCancel()"
      @ok="handleUpdateRoleSubmit()"
    >
      <a-form
        ref="FormUpdateRoleRef"
        :model="formUpdateRole"
        :rules="formUpdateRoleRules"
        :label-col="labelCol"
        :wrapper-col="wrapperCol"
      >
        <a-form-item label="名称:" name="name">
          <a-input v-model:value="formUpdateRole.name" placeholder="请输入角色名称" />
        </a-form-item>
        <a-form-item label="标识:" name="slug">
          <a-input v-model:value="formUpdateRole.slug" placeholder="请输入角色标识" disabled />
        </a-form-item>
        <a-form-item label="描述:" name="desc">
          <a-input v-model:value="formUpdateRole.desc" placeholder="请输入角色描述" />
        </a-form-item>
        <a-form-item label="角色:" name="parentRole">
          <a-select disabled v-model:value="formUpdateRole.parentId" :options="options"></a-select>
        </a-form-item>
      </a-form>
    </Modal>
  </div>
</template>

<script setup lang="ts">
  import fetchApi from '/@/api/role';
  import { ChannelStatusList } from '/@/views/channel/constant';
  import { AuthEnum } from '/@/enums/authEnum';
  import { useMessage } from '/@/hooks/useMessage';
  import { FormInstance } from 'ant-design-vue';
  import { Ref } from 'vue/dist/vue';
  import { Rule } from 'ant-design-vue/es/form';
  import type { SelectProps } from 'ant-design-vue';
  import { RoleCreate, RoleUpdate } from '/@/api/role/model';
  import { ref } from 'vue';

  const { createMessage } = useMessage();
  const ELRef = ref<{ refresh: () => void }>();
  const refresh = () => ELRef.value?.refresh();

  const labelCol = { style: { width: '110px' } };
  const wrapperCol = { span: 17 };

  const options = ref<SelectProps['options']>([]);

  const tableFilterButton = reactive({
    type: 'primary',
    label: '新增角色',
    auth: AuthEnum.role_create,
    onClick: () => {
      modalCreateRole.visible = true;
      options.value = [
        {
          value: '0',
          label: '根角色',
        },
      ];
    },
  });

  const modalCreateRole = reactive({
    loading: false,
    visible: false,
    title: '创建角色',
    okText: '创建',
  });

  const formCreateRole: Ref<RoleCreate> = ref({
    name: '',
    slug: '',
    desc: '',
    parentId: 0,
  });

  const modalUpdateRole = reactive({
    loading: false,
    visible: false,
    title: '编辑角色',
    okText: '更新',
  });

  const formUpdateRole: Ref<RoleUpdate> = ref({
    id: -1,
    name: '',
    slug: '',
    desc: '',
    parentId: 0,
  });

  const checkSlug = async (_rule: Rule, value: string) => {
    if (!value) {
      return Promise.reject('请输入角色标识');
    }
    if (value.length < 3 || value.length > 6) {
      return Promise.reject('标识符需要在3-6位');
    } else {
      return Promise.resolve();
    }
  };

  const checkDesc = async (_rule: Rule, value: string) => {
    if (value.length > 20) {
      return Promise.reject('标识符需要小于20位');
    } else {
      return Promise.resolve();
    }
  };

  const formCreateRoleRules: Record<string, Rule[]> = {
    name: [{ required: true, trigger: 'blur', message: '请输入角色名称' }],
    slug: [{ required: true, validator: checkSlug, trigger: 'blur' }],
    desc: [{ validator: checkDesc, trigger: 'blur' }],
  };

  const formUpdateRoleRules: Record<string, Rule[]> = {
    name: [{ required: true, trigger: 'blur', message: '请输入角色名称' }],
    desc: [{ validator: checkDesc, trigger: 'blur' }],
  };

  const FormCreateRoleRef = ref<FormInstance>();
  const FormUpdateRoleRef = ref<FormInstance>();
  const handelCreateRoleCancel = () => {
    modalCreateRole.visible = false;
    FormCreateRoleRef.value?.resetFields();
  };

  const handelUpdateRoleCancel = () => {
    modalUpdateRole.visible = false;
    FormUpdateRoleRef.value?.resetFields();
  };

  const handleCreateRoleSubmit = () => {
    FormCreateRoleRef.value?.validate().then(async () => {
      modalCreateRole.loading = true;
      const res = await fetchApi.CreateRole(formCreateRole.value);
      modalCreateRole.loading = false;
      if (res) {
        modalCreateRole.visible = false;
        createMessage.success('创建角色成功');
        handelCreateRoleCancel();
        refresh();
      } else {
        createMessage.error('创建角色失败');
      }
    });
  };

  const handleUpdateRoleSubmit = () => {
    FormUpdateRoleRef.value?.validate().then(async () => {
      modalUpdateRole.loading = true;
      const res = await fetchApi.UpdateRole(formUpdateRole.value);
      modalUpdateRole.loading = false;
      if (res) {
        modalUpdateRole.visible = false;
        createMessage.success('更新角色成功');
        handelUpdateRoleCancel();
        refresh();
      } else {
        createMessage.error('更新角色失败');
      }
    });
  };

  const tableActions = reactive([
    {
      label: '编辑权限',
      auth: AuthEnum.channel_show,
    },
    {
      label: '新增子角色',
      auth: AuthEnum.channel_show,
      onClick: async (row) => {
        modalCreateRole.visible = true;
        formCreateRole.value.parentId = row.id;
        options.value = [
          {
            value: row.id,
            label: row.name,
          },
        ];
      },
    },
    {
      label: '编辑',
      auth: AuthEnum.device_update,
      onClick: async (row) => {
        modalUpdateRole.visible = true;
        formUpdateRole.value.id = row.id;
        formUpdateRole.value.slug = row.slug;
        formUpdateRole.value.name = row.name;
        formUpdateRole.value.desc = row.desc;
        options.value = [
          {
            value: row.id,
            label: row.name,
          },
        ];
      },
    },
    {
      label: '删除',
      popConfirm: {
        title: '确认删除吗？',
        onConfirm: async (row) => {
          const result = await fetchApi.DeleteRole(row.id);
          if (result) {
            refresh()
            createMessage.success('删除成功');
          }
        },
      },
      auth: AuthEnum.device_delete,
    },
  ]);

  const columns = [
    {
      title: '名称',
      dataIndex: 'name',
      key: 'name',
      width: '20%',
    },
    {
      title: '标识',
      dataIndex: 'slug',
      key: 'slug',
      width: '20%',
    },
    {
      title: '描述',
      dataIndex: 'desc',
      key: 'desc',
      width: '20%',
    },
    {
      title: '操作',
      dataIndex: 'action',
      width: '40%',
      key: 'action',
    },
  ];
</script>

<style scoped lang="less"></style>
