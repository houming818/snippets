# -*- coding: utf-8 -*-

import uuid

from django.db import models

class CoreModelManager(models.Manager):

    def get_queryset(self):
        # 过滤掉已删除的数据
        return super().get_queryset().filter(deleted=0)

    def create_uuid(self, *args, **kwargs):
        # 创建一个新的 UUID
        new_uuid = uuid.uuid4()

        # 使用 UUID 创建对象
        return self.create(id=new_uuid, *args, **kwargs)

    def delete(self, *args, **kwargs):
        # 软删除，将 is_deleted 设置为 True
        self.deleted = 1
        self.save()

    def hard_delete(self):
        # 硬删除，使用原始的 delete 方法
        super().delete()

# 新建一个baseModel类,继承于django.db.models，用于继承
class CoreModel(models.Model):

    # 新建一个字段，字段名为id，类型为uuid，主键，不可编辑，自动生成
    id = models.UUIDField(primary_key=True, default=uuid.uuid4, editable=False)

    # 新建一个字段，字段名为last_update，类型为datetime，自动更新，最后更新时间
    last_update = models.DateTimeField(auto_now=True, verbose_name='最后更新时间')

    # 新建一个字段，字段名为created_time，类型为datetime，自动创建，创建时间
    created_time = models.DateTimeField(auto_now_add=True, verbose_name='创建时间')

    # 新建一个字段，字段名为deleted，类型为bool，是否删除
    deleted = models.IntegerField(default=0, verbose_name='是否删除')

    # 自定义管理器
    objects = CoreModelManager()

    class Meta:
        abstract = True
