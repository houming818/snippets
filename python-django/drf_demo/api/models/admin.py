
from django.db import models
from . import CoreModel

# 新建一个Modle类，继承于CoreModel,类名AdminConfig
class AdminConfig(CoreModel):
    # 新建一个字段，字段名为name，类型为char，最大长度为50，唯一
    name = models.CharField(max_length=50, unique=True, verbose_name='配置名称')
    # 新建一个字段，字段名为value，类型为text，唯一
    value = models.TextField(verbose_name='配置值')

    class meta:
        verbose_name = '系统配置'
        verbose_name_plural = verbose_name
