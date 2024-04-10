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
  video_play = 'video:play',
  video_record = `video:record`,
  video_stop_record = 'video:stop:record',
  channel_update = 'channel:update',
  channel_delete = 'channel:delete',

  /**
   * 角色管理
   */
  role_create = 'role:create',
  role_update = 'role:update',
  role_delete = 'role:delete',
  role_allPermission = 'role:allPermission',

  /**
   * 远端录像
   */
  record_range = 'record:range',
  record_details = 'record:details',
  record_play = 'record:play',

  /**
   * 用户管理
   */
  user_create = 'user:create',
  user_enable = 'user:enable',
  user_update = 'user:update',
  user_delete = 'user:delete',
}
