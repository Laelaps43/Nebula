<template>
  <div class="node-conatiner">
    <h2 class="nc_title font18">用户管理</h2>
    <Table
      ref="ELRef"
      :url="fetchApi.user_page"
      :columns="columns"
      :hiddenFilter="false"
      :button="tableFilterButton"
      :actions="tableActions"
      :status-list="enableStatusList"
      :role-list="roleList"
    />
    <!-- 新建用户 -->
    <Modal v-bind="modalCreateUser" @cancel="handleCreateUserCancel" @ok="handleCreateUserSubmit">
      <a-form
        ref="FormRef"
        :model="formCreateUser"
        :rules="rules"
        :label-col="labelCol"
        :wrapper-col="wrapperCol"
      >
        <a-form-item label="名称:" name="name">
          <a-input v-model:value="formCreateUser.name" placeholder="请输入用户名" />
        </a-form-item>
        <a-form-item label="邮箱:" name="email">
          <a-input v-model:value="formCreateUser.email" placeholder="请输入邮箱" />
        </a-form-item>
        <a-form-item label="角色:" name="role">
          <a-cascader
            v-model:value="userRoleSelectCreateUser"
            style="width: 100%"
            multiple
            max-tag-count="responsive"
            :allowClear="false"
            :options="roleList"
            placeholder="请选择角色"
            :show-checked-strategy="Cascader.SHOW_CHILD"
          />
        </a-form-item>
        <a-form-item label="密码:" name="password">
          <a-input
            v-model:value="formCreateUser.password"
            type="password"
            placeholder="请输入密码"
          />
        </a-form-item>
        <a-form-item label="密码:" name="passwordAgain">
          <a-input
            v-model:value="formCreateUser.passwordAgain"
            type="password"
            placeholder="请在输入一次密码"
          />
        </a-form-item>
      </a-form>
    </Modal>
    <!-- 编辑用户 -->
    <Modal v-bind="modalUpdateUser" @cancel="handleUpdateUserCancel" @ok="handleUpdateUserSubmit">
      <a-form
        ref="FormRef"
        :model="formUpdateUser"
        :rules="rulesUpdate"
        :label-col="labelCol"
        :wrapper-col="wrapperCol"
      >
        <a-form-item label="名称:" name="name">
          <a-input v-model:value="formUpdateUser.name" placeholder="请输入用户名" />
        </a-form-item>
        <a-form-item label="邮箱:" name="email">
          <a-input v-model:value="formUpdateUser.email" placeholder="请输入邮箱" />
        </a-form-item>
        <a-form-item label="角色:" name="role">
          <a-cascader
            v-model:value="userRoleSelectCreateUser"
            style="width: 100%"
            multiple
            max-tag-count="responsive"
            :allowClear="false"
            :options="roleList"
            placeholder="请选择角色"
            :show-checked-strategy="Cascader.SHOW_CHILD"
          />
        </a-form-item>
      </a-form>
    </Modal>
  </div>
</template>
<script setup lang="ts">
  import { columns, enableStatusList } from './constant';
  import { useMessage } from '/@/hooks/useMessage';
  import fetchApi from '/@/api/user';
  import fetchRoleApi from '/@/api/role';
  import type { FormInstance } from 'ant-design-vue';
  import type { Rule } from 'ant-design-vue/es/form';
  import { reactive, ref, Ref } from 'vue';
  import { AuthEnum } from '/@/enums/authEnum';
  import { CreateUser, UpdateUser } from "/@/api/user/model";
  import { Cascader } from 'ant-design-vue';

  const { createMessage } = useMessage();
  const FormRef = ref<FormInstance>();
  const ELRef = ref<{ refresh: () => void }>();
  const refresh = () => ELRef.value?.refresh();

  const labelCol = { style: { width: '110px' } };
  const wrapperCol = { span: 17 };

  const tableFilterButton = reactive({
    type: 'primary',
    label: '新增用户',
    auth: AuthEnum.device_create,
    onClick: () => {
      modalCreateUser.visible = true;
    },
  });

  // 新建用户
  const modalCreateUser = reactive({
    loading: false,
    visible: false,
    title: '新建用户',
    okText: '新建',
  });
  // 新建用户
  const modalUpdateUser = reactive({
    loading: false,
    visible: false,
    title: '编辑用户',
    okText: '更新',
  });
  const userRoleSelectCreateUser = ref<number[]>([]);

  const validatePass = async (_rule: Rule, value: string) => {
    if (value === '') {
      return Promise.reject('请输入密码');
    } else {
      if (formCreateUser.value.password !== '') {
        FormRef.value.validateFields('passwordAgain');
      }
      return Promise.resolve();
    }
  };
  const validatePassAgain = async (_rule: Rule, value: string) => {
    if (value === '') {
      return Promise.reject('请再一次输入密码');
    } else if (value !== formCreateUser.value.password) {
      return Promise.reject('两次输入的密码不一致');
    } else {
      return Promise.resolve();
    }
  };
  const validateEmail = async (_rule: Rule, value: string) => {
    const emailPattern = /^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$/;
    if (value === '') {
      return Promise.reject('请输入邮箱');
    } else if (!emailPattern.test(value)) {
      return Promise.reject('请输入正确的邮箱');
    } else {
      return Promise.resolve();
    }
  };
  const validateRole = async (_rule: Rule, value: string) => {
    if (value === '') {
      return Promise.reject('请选择角色');
    } else if (userRoleSelectCreateUser.value.length === 0) {
      return Promise.reject('请选择角色');
    } else {
      return Promise.resolve();
    }
  };

  // form
  const rules: Record<string, Rule[]> = {
    name: [{ required: true, trigger: 'blur', message: '请输入用户名' }],
    email: [{ required: true, trigger: 'blur', validator: validateEmail }],
    role: [{ required: true, trigger: 'blur', validator: validateRole }],
    password: [{ required: true, validator: validatePass, trigger: 'change' }],
    passwordAgain: [{ required: true, validator: validatePassAgain, trigger: 'change' }],
  };
  const rulesUpdate: Record<string, Rule[]> = {
    name: [{ required: true, trigger: 'blur', message: '请输入用户名' }],
    email: [{ required: true, trigger: 'blur', validator: validateEmail }],
    role: [{ required: true, trigger: 'blur', validator: validateRole }],
  };

  const formUpdateUser: Ref<UpdateUser> = ref({
    id: '',
    name: '',
    email: '',
    roles: [],
  });

  const formCreateUser: Ref<CreateUser> = ref({
    name: '',
    email: '',
    password: '',
    passwordAgain: '',
    roles: [],
  });

  const tableActions = reactive([
    {
      label: '冻结',
      onClick: async (row) => {
        let res = await fetchApi.user_enable({ id: row.id, enable: row.enable ? 0 : 1 });
        if (res) {
          createMessage.success('更新成功');
          refresh();
        }
      },
      auth: AuthEnum.channel_show,
    },
    {
      label: '编辑',
      auth: AuthEnum.device_update,
      onClick: async (row) => {
        modalUpdateUser.visible = true;
        formUpdateUser.value.id = row.id;
        formUpdateUser.value.email = row.email;
        formUpdateUser.value.name = row.username;
        console.log(formUpdateUser.value);
        userRoleSelectCreateUser.value = row.roles?.map((role) => [role?.id]).filter(Boolean) || [];
      },
    },
    {
      label: '删除',
      popConfirm: {
        title: '确认删除吗？',
        onConfirm: async (row) => {
          const res = await fetchApi.user_delete(row.id);
          if (res) {
            createMessage.success('删除用户成功');
            refresh();
          }
        },
      },
      auth: AuthEnum.device_delete,
    },
  ]);

  const handleCancel = () => {
    FormRef.value?.resetFields();
  };

  const handleCreateUserCancel = () => {
    modalCreateUser.visible = false;
    FormRef.value?.resetFields();
  };

  const handleUpdateUserCancel = () => {
    modalUpdateUser.visible = false;
    FormRef.value?.resetFields();
  };

  const handleCreateUserSubmit = () => {
    FormRef.value?.validate().then(async () => {
      modalCreateUser.loading = true;
      formCreateUser.value.roles = userRoleSelectCreateUser.value.flatMap((list) => list || []);
      const res = await fetchApi.user_create(formCreateUser.value);
      modalCreateUser.loading = false;
      modalCreateUser.visible = false;
      if (res) {
        createMessage.success('创建用户成功');
        userRoleSelectCreateUser.value = [];
        handleCancel();
        refresh();
      } else {
        createMessage.error('创建用户失败');
      }
    });
  };
  const handleUpdateUserSubmit = () => {
    FormRef.value?.validate().then(async () => {
      modalUpdateUser.loading = true;
      formUpdateUser.value.roles = userRoleSelectCreateUser.value.flatMap((list) => list || []);
      const res = await fetchApi.user_update(formUpdateUser.value);
      modalUpdateUser.loading = false;
      modalUpdateUser.visible = false;
      if (res) {
        createMessage.success('编辑用户成功');
        userRoleSelectCreateUser.value = [];
        handleCancel();
        refresh();
      } else {
        createMessage.error('编辑用户失败');
      }
    });
  };

  const roleList = ref<
    {
      label: string;
      value: string;
    }[]
  >([]);
  const getRoleList = async () => {
    const res = await fetchRoleApi.ListRole({ limit: 999, page: 1 });
    if (res) {
      res.list.forEach((role) => {
        roleList.value?.push({
          label: role.name,
          value: role.id,
        });
      });
    } else {
      createMessage.error('获取角色失败');
    }
  };
  onMounted(() => {
    getRoleList();
  });
</script>
<style lang="less" scoped>
  .node-conatiner {
    .nc_title {
      margin-top: 6px;
      margin-bottom: 30px;
    }
  }
</style>
