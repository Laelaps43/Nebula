/**
 * @name AuthEnum
 * @description 权限，配合指令 v-auth 使用
 * @Example v-auth="AuthEnum.user_create"
 */

export enum AuthEnum {
  /**
   * 设备列表
   */
  // 编辑设备
  device_update = 'device:update',
  // 删除设备
  device_delete = 'device:delete',
  // 通道展示
  channel_show = 'device:showChannel',
  // 增添设备
  device_create = 'device:create',

  /**
   * 通道列表
   */
  channel_create = 'channel:create',

  /**
   * 角色管理
   */
  role_create = 'role:create',
}
