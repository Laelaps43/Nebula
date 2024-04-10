<template>
  <div class="sys-setting">
    <a-dropdown placement="bottom">
      <template #overlay>
        <a-menu :selectedKeys="selectedKeys" class="menu-box">
          <a-menu-item v-for="item in navs" :key="item.type" @click="handleRoute(item)">
            <template #icon>
              <Icon align="1px" size="20px" :type="item.icon" />
            </template>
            <span>{{ item.name }}</span>
          </a-menu-item>
        </a-menu>
      </template>
      <Space class="wrap" align="baseline" direction="horizontal">
        <span class="setting">我的</span>
        <Icon align="2px" type="xialajiantouxiao" />
      </Space>
    </a-dropdown>
    <Modal
      v-bind="modalRoleSwitch"
      @cancel="handelSwitchRoleCancel()"
      @ok="handleSwitchRoleSubmit()"
    >
      <div style="width: 90%">
        <a-radio-group v-model:value="formSwitchRole" name="radioGroup" style="width: 100%">
          <div style="width: 100%; display: flex; justify-content: center; flex-direction: column">
            <a-row v-for="(item, index) in roleOptions" :key="index">
              <a-col v-for="(col, y) in item" :span="6" :key="y">
                <a-radio :value="col.value"
                  ><span :style="{ color: col.hold ? 'red' : 'black' }">{{
                    col.label
                  }}</span></a-radio
                >
              </a-col>
            </a-row>
          </div>
        </a-radio-group>
      </div>
    </Modal>
  </div>
</template>

<script setup lang="ts">
  import { Space } from 'ant-design-vue';
  import { useUserStore } from '/@/store/modules/user';
  import { navs as myNavs } from './constant';
  import { usePermissioStore } from '/@/store/modules/permission';
  import { useMessage } from '/@/hooks/useMessage';
  import fetchApi from '/@/api/user';

  const store = useUserStore();
  const permissioStore = usePermissioStore();
  const router = useRouter();

  const navs = ref(myNavs);
  const selectedKeys = ref<string[]>([]);

  watchEffect(() => {
    const modules = permissioStore.getModules;
    if (modules.length) {
      navs.value = unref(navs).filter((n) => (n.auth ? modules.includes(n.auth) : true));
    }
  });

  watchEffect(() => {
    if (router.currentRoute) {
      const matched = router.currentRoute.value.matched.concat();
      selectedKeys.value = matched.filter((r) => r.name !== 'index').map((r) => r.path);
    }
  });

  store.fetchUserRole();

  const modalRoleSwitch = reactive({
    loading: false,
    visible: false,
    title: '选择角色',
    okText: '确认',
  });

  const roleOptions = ref<any[][]>([]);

  const { createMessage } = useMessage();
  const handelSwitchRoleCancel = () => {
    modalRoleSwitch.visible = false;
  };

  const formSwitchRole = ref<String>('1');

  const handleSwitchRoleSubmit = async () => {
    modalRoleSwitch.loading = true;
    let res = await fetchApi.user_switch({ roleId: formSwitchRole.value });
    store.switchRole();
    modalRoleSwitch.loading = false;
    if (res) {
      createMessage.success('设置权限成功');
    } else {
      createMessage.error('设置权限失败失败');
    }
  };

  const convertToTwoDimensionalArray = (inputArray: any[]) => {
    const result: any[][] = [];

    for (let i = 0; i < inputArray.length; i += 4) {
      result.push(inputArray.slice(i, i + 4));
    }
    return result;
  };
  const handleRoute = (item?: any) => {
    if (item.type === 'myAccount') return router.push(item.path);
    else if (item.type === 'switchRole') {
      modalRoleSwitch.visible = true;
      let roleSet = store.getRoleSet();
      roleOptions.value = convertToTwoDimensionalArray(roleSet);
      formSwitchRole.value = roleSet.find((item) => item.hold == true).value;
    } else {
      store.logout();
    }
    // 退出登录
  };
</script>

<style lang="less" scoped>
  .sys-setting {
    height: 100%;
    display: flex;
    justify-content: center;
    padding-right: 16px;

    .wrap {
      display: flex;
      height: 55px;

      .setting {
        font-size: 16px;
        font-weight: 600;
        line-height: 22px;
        color: rgba(0, 0, 0, 0.85);
        margin: 0 0px 0 4px;
      }
    }

    .my-icon {
      font-size: 18px !important;
    }
  }

  .menu-box :deep(.ant-dropdown-menu-item) {
    width: 142px;
    height: 42px;
    line-height: 42px;
    padding: 0 16px;
  }

  .menu-box :deep(.ant-dropdown-menu-item-selected) {
    background: #eaeffe;
    color: #3860f4;
  }
</style>
