from typing import Tuple

# pyright:  reportMissingImports=false
from api.services import CoreService

from api.models import AdminConfig


class AdminService(CoreService):
    admin_config = AdminConfig

    def config(self, config_name: str)-> Tuple[dict, str]:
        """
        获取配置
        :param config_name: 配置名称
        :return: 配置
        """
        config = self.admin_config.objects.filter(name=config_name).first()
        if config:
            return config.value, ''
        return {}, '配置不存在'
