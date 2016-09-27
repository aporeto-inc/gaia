# -*- coding: utf-8 -*-



<<<<<<< 5c555e817c0bd485e9181ed3ec6ccc194c606dac
__all__ = ['APIAuthorizationPolicy', 'Authenticator', 'Certificate', 'ClairNotification', 'ComputedDependencyMapView', 'ComputedPolicy', 'DependencyMap', 'DependencyMapSubview', 'DependencyMapView', 'ExternalService', 'FileAccessPolicy', 'FilePath', 'FlowStatistic', 'HealthReport', 'Image', 'Layer', 'MapEdge', 'MapNode', 'MyNamespace', 'Namespace', 'NamespaceMappingPolicy', 'NetworkAccessPolicy', 'Notification', 'Policy', 'PolicyRule', 'ProcessingUnit', 'RenderedPolicy', 'Root', 'Server', 'Service', 'SystemCall', 'Tag', 'User', 'Vulnerability', 'VulnerabilityWithLayers', 'routes']
=======
__all__ = ['APIAuthorizationPolicy', 'Authenticator', 'Certificate', 'ClairNotification', 'DependencyMap', 'DependencyMapSubview', 'DependencyMapView', 'ExternalService', 'FileAccessPolicy', 'FilePath', 'FlowStatistic', 'HealthReport', 'Image', 'Integration', 'Layer', 'MapEdge', 'MapNode', 'MyNamespace', 'Namespace', 'NamespaceMappingPolicy', 'NetworkAccessPolicy', 'Notification', 'Policy', 'PolicyRule', 'ProcessingUnit', 'RenderedPolicy', 'Root', 'Server', 'SystemCall', 'Tag', 'User', 'Vulnerability', 'VulnerabilityWithLayers', 'routes']
>>>>>>> Fixed: renamed service to integration

from .apiauthorizationpolicy import APIAuthorizationPolicy
from .authenticator import Authenticator
from .certificate import Certificate
from .clairnotification import ClairNotification
from .computeddependencymapview import ComputedDependencyMapView
from .computedpolicy import ComputedPolicy
from .dependencymap import DependencyMap
from .dependencymapsubview import DependencyMapSubview
from .dependencymapview import DependencyMapView
from .externalservice import ExternalService
from .fileaccesspolicy import FileAccessPolicy
from .filepath import FilePath
from .flowstatistic import FlowStatistic
from .healthreport import HealthReport
from .image import Image
from .integration import Integration
from .layer import Layer
from .mapedge import MapEdge
from .mapnode import MapNode
from .mynamespace import MyNamespace
from .namespace import Namespace
from .namespacemappingpolicy import NamespaceMappingPolicy
from .networkaccesspolicy import NetworkAccessPolicy
from .notification import Notification
from .policy import Policy
from .policyrule import PolicyRule
from .processingunit import ProcessingUnit
from .renderedpolicy import RenderedPolicy
from .root import Root
from .server import Server
from .systemcall import SystemCall
from .tag import Tag
from .user import User
from .vulnerability import Vulnerability
from .vulnerabilitywithlayers import VulnerabilityWithLayers
from .routes import routes