<template>
  <a-spin :spinning="loading">
    <h2 class="font18 marT13 rowSC link" @click="handleGoBack">
      <LeftOutlined />
      <span class="marL10">{{ '返回' }}</span>
    </h2>
    <a-divider />
    <div class="title font18">用户信息</div>
    <div class="content">
      <div class="content_left">
        <!-- <div class="img"><img src="" alt="" /></div> -->
        <a-avatar :size="100" :src="avatar" />
      </div>
      <div class="content_right">
        <a-row class="row" v-for="item in columns" :key="item.label">
          <a-col :span="12">
            <a-space>
              <div class="label-w">{{ item.label }}</div>
              <div v-if="modifyItem.key === item.key">
                <a-input v-model:value="modifyItem.value" />
              </div>
              <div v-else-if="item.key == 'role'">{{ item.value.name }}</div>
              <div v-else>{{ item.value }}</div>
            </a-space>
          </a-col>
          <a-col :span="12">
            <a-space v-if="modifyItem.key === item.key">
              <a @click="handleSubmit">保存</a>
              <a @click="handleCancle">取消</a>
            </a-space>
            <div v-else>
              <a-button v-if="item.isEdit" type="link" @click="handleClick(item)"> 修改</a-button>
            </div>
          </a-col>
        </a-row>
      </div>
    </div>
  </a-spin>
</template>
<script setup lang="ts">
  import avatar from '/@/assets/images/avatar.png';
  import { KeyValue } from './constant';
  import { useSysAccountStore } from '/@/store/modules/sysAccount';
  import { useMessage } from '/@/hooks/useMessage';
  import { LeftOutlined } from '@ant-design/icons-vue';

  const { createMessage } = useMessage();
  const store = useSysAccountStore();

  const router = useRouter();

  const handleGoBack = () => {
    router.back();
  };
  const initVal = {
    key: '',
    value: '',
  };

  const loading = ref(false);
  const columns = ref(KeyValue);
  const modifyItem = ref(initVal);
  const account = computed(() => store.getAccount);

  watch(
    () => account.value,
    (val) => {
      columns.value = KeyValue.map((n) => ({
        ...n,
        value: toRaw(val)?.[n.key] || n.value,
      }));
    },
  );

  onMounted(() => store.fetchAccount());

  onUnmounted(() => store.resetState());

  //
  const handleClick = (item) => {
    modifyItem.value = item.key === 'password' ? { ...item, value: '' } : item;
  };

  const handleCancle = () => {
    modifyItem.value = initVal;
  };

  const handleSubmit = async () => {
    const { key, value } = unref(modifyItem);
    const params = { id: account.value!.id, [key]: value };
    //
    loading.value = true;
    let res = await store.fetchAccountUpdate(params);
    if (res) {
      createMessage.success(`修改成功`);
    } else {
      createMessage.error(`修改失败`);
    }
    loading.value = false;
    handleCancle();
  };
</script>
<style lang="less" scoped>
  .title {
    margin-top: -24px;
    height: 105px;
    line-height: 105px;
    padding-left: 29px;
  }

  .content {
    display: flex;

    &_left {
      padding-left: 15px;
      margin-right: 24px;
      // .img {
      //   width: 100px;
      //   height: 100px;
      //   border-radius: 50%;
      //   img {
      //     width: 100%;
      //     height: 100%;
      //   }
      // }
    }

    &_right {
      flex: 1;
      padding-bottom: 44px;

      .row {
        height: 53px;
        padding-left: 24px;
        line-height: 53px;
        font-size: 14px;
        color: #333333;

        .label-w {
          width: 80px;
        }

        &:nth-child(odd) {
          background-color: #fafafa;
        }
      }
    }
  }
</style>
