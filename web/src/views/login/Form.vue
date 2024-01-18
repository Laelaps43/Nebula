<template>
  <div class="form_box">
    <a-form :model="formModel" :rules="rules" @finish="handleFinish">
      <p class="text">请输入邮箱</p>
      <a-form-item name="email">
        <a-input
          class="reset-input"
          v-model:value="formModel.email"
          placeholder="邮箱"
        >
          <template #prefix>
            <!-- <user-outlined class="icon" type="user" /> -->
            <Icon size="24px" type="youxiangdenglu" class="icon" />
          </template>
        </a-input>
      </a-form-item>
      <p class="text">请输入密码</p>
      <a-form-item name="password">
        <a-input
          class="reset-input"
          v-model:value="formModel.password"
          type="password"
          placeholder="密码">
          <template #prefix>
            <!-- <lock-outlined class="icon" /> -->
            <Icon size="22px" type="iconqingshurumimadengluye" class="icon" />
          </template>
        </a-input>
      </a-form-item>
      <a-form-item>
        <a-button html-type="submit" class="btn" :loading="loading">立即登录</a-button>
      </a-form-item>
    </a-form>
  </div>
</template>
<script setup lang="ts">
  import { useUserStore } from '/@/store/modules/user';

  const userStore = useUserStore();
  const router = useRouter();

  const loading = ref(false);

  let state: any = reactive({
    otherQuery: {},
    redirect: undefined,
  });

  /* listen router change  */
  const route = useRoute();
  let getOtherQuery = (query: any) => {
    return Object.keys(query).reduce((acc: any, cur) => {
      if (cur !== 'redirect') {
        acc[cur] = query[cur];
      }
      return acc;
    }, {});
  };

  watch(
    route,
    (route) => {
      const query = route.query;
      if (query) {
        state.redirect = query.redirect;
        state.otherQuery = getOtherQuery(query);
      }
    },
    { immediate: true },
  );

  const rules = {
    email: [{ required: true, trigger: 'blur', message: '请输入邮箱' }],
    password: [{ required: true, trigger: 'blur', message: '请输入密码' }],
  };

  const formModel = reactive({
    email: '',
    password: '',
  });

  const handleFinish = async (values) => {
    loading.value = true;
    const res = await userStore.login(values);
    loading.value = false;
    if (res) {
      // message.success('成功');
      // router.replace({ path: state.redirect || '/', query: state.otherQuery });
      router.replace('/');
    }
  };
</script>
<style lang="less">
  .form_box {
    margin-top: 30px;
    .btn {
      width: 100%;
      height: 54px;
      background: linear-gradient(90deg, #00c3fd 0%, #3662f4 100%);
      border-radius: 6px;
      color: #ffffff;
      font-size: 20px;
      &:hover {
        opacity: 0.85;
        border: none;
      }
    }
    .icon {
      color: #666666;
    }
    .text {
      font-size: 14px;
      line-height: 20px;
      color: #999999;
      letter-spacing: 1.1px;
      margin-bottom: 10px;
    }
    .gray_text {
      font-size: 12px;
      color: #666666;
    }
    .reset_checkbox {
      .ant-checkbox-inner {
        border-radius: 50%;
      }
      & > span:last-child {
        font-size: 12px;
        color: #666666;
      }
    }
    .reset-input {
      height: 50px;
      line-height: 50px;
      border: 1px solid #707070;
      border-radius: 6px;
    }
    .copyright {
      margin-top: 20px;
      font-size: 12px;
      color: #999999;
      opacity: 0.85;
    }
  }
</style>
