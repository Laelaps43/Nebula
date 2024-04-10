<template>
  <Modal v-bind="modalPermission" @cancel="handelPermissionCancel()" @ok="handleCreateRoleSubmit()">
    <a-form
      ref="FormPermissionRef"
      :model="formPermission"
      :label-col="labelCol"
      :wrapper-col="wrapperCol"
    />
  </Modal>
</template>
<script lang="ts" setup>
  import { ref } from 'vue/dist/vue';
  import { FormInstance } from 'ant-design-vue';
  import { Ref } from 'vue';
  import { RoleCreate } from '/@/api/role/model';
  import { useMessage } from '/@/hooks/useMessage';

  const modalPermission = reactive({
    loading: false,
    visible: false,
    title: '设置权限',
    okText: '确认',
  });

  const showDialog = () => {
    modalPermission.visible = true;
  };

  const FormPermissionRef = ref<FormInstance>();

  const labelCol = { style: { width: '110px' } };
  const wrapperCol = { span: 17 };
  const handelPermissionCancel = () => {
    modalPermission.visible = false;
    FormPermissionRef.value?.resetFields();
  };

  const formPermission: Ref<RoleCreate> = ref({
    name: '',
    slug: '',
    desc: '',
    parentId: 0,
  });

  const { createMessage } = useMessage();

  const handleCreateRoleSubmit = () => {
    FormPermissionRef.value?.validate().then(async () => {
      modalPermission.loading = true;
      modalPermission.loading = false;
      // if (res) {
      modalPermission.visible = false;
      createMessage.success('创建角色成功');
      handelPermissionCancel();
      // refresh();
      // } else {
      //   createMessage.error('创建角色失败');
      // }
      // });
    });
  };
</script>
