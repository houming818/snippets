# Written by: HouMing

# pyright:  reportMissingImports=false
from .core import CoreService
from .admin import AdminService

__all__ = ['CoreService', 'AdminService', 'CoreFactory']

class CoreFactory(object):

    @classmethod
    def get_admin_service(cls, *args, **kwargs) -> AdminService:
        return AdminService(*args, **kwargs)
