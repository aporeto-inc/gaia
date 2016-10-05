# -*- coding: utf-8 -*-

from pyelemental import RESTObject
from pyelemental import validate_string_in_list, validate_float_in_list, validate_int_in_list, validate_required_int, validate_required_float, validate_required_string, validate_required_time, validate_maximum_float, validate_minimum_float, validate_maximum_int, validate_minimum_int, validate_maximum_length, validate_minimum_length, validate_pattern


class SystemInfo(RESTObject):
    """ Represents a SystemInfo in the 

        Notes:
            A SystemInfo contains the status and various information about the Server.
    """

    def __init__(self, **kwargs):
        """ Initializes a SystemInfo instance

          Notes:
              You can specify all parameters while calling this methods.
              A special argument named `data` will enable you to load the
              object from a Python dictionary

          Examples:
              >>> systeminfo = SystemInfo(id=u'xxxx-xxx-xxx-xxx', name=u'SystemInfo')
              >>> systeminfo = SystemInfo(data=my_dict)
        """

        super(SystemInfo, self).__init__()

        # Read/Write Attributes
        
        self._apiversion = None
        self._bahamutversion = None
        self._elementalversion = None
        self._gaiaversion = None
        self._manipulateversion = None
        self._midgardurl = None
        self._squallversion = None
        self._status = None
        
        self.expose_attribute(local_name="APIVersion", remote_name="APIVersion")
        self.expose_attribute(local_name="bahamutVersion", remote_name="bahamutVersion")
        self.expose_attribute(local_name="elementalVersion", remote_name="elementalVersion")
        self.expose_attribute(local_name="gaiaVersion", remote_name="gaiaVersion")
        self.expose_attribute(local_name="manipulateVersion", remote_name="manipulateVersion")
        self.expose_attribute(local_name="midgardURL", remote_name="midgardURL")
        self.expose_attribute(local_name="squallVersion", remote_name="squallVersion")
        self.expose_attribute(local_name="status", remote_name="status")

        self._compute_args(**kwargs)

    def __str__(self):
        return '<%s:%s>' % (self.identity()["name"], self.identifier)

    def identifier(self):
        """ Identifier returns the value of the object's unique identifier.
        """
        return self.

    def setIdentifier(self, ID):
        """ SetIdentifier sets the value of the object's unique identifier.
        """
        self. = ID

    def identity(self):
        """ Identity returns the Identity of the object.
        """
        return systeminfoIdentity

    # Properties
    @property
    def APIVersion(self):
        """ Get APIVersion value.

          Notes:
              APIVersion is the API version served by the server.

              
        """
        return self._apiversion

    @APIVersion.setter
    def APIVersion(self, value):
        """ Set APIVersion value.

          Notes:
              APIVersion is the API version served by the server.

              
        """
        self._apiversion = value
    
    @property
    def bahamutVersion(self):
        """ Get bahamutVersion value.

          Notes:
              BahamutVersion is the version of Bahamut used by the server.

              
        """
        return self._bahamutversion

    @bahamutVersion.setter
    def bahamutVersion(self, value):
        """ Set bahamutVersion value.

          Notes:
              BahamutVersion is the version of Bahamut used by the server.

              
        """
        self._bahamutversion = value
    
    @property
    def elementalVersion(self):
        """ Get elementalVersion value.

          Notes:
              ElementalVersion is the version of Elemental used by the server.

              
        """
        return self._elementalversion

    @elementalVersion.setter
    def elementalVersion(self, value):
        """ Set elementalVersion value.

          Notes:
              ElementalVersion is the version of Elemental used by the server.

              
        """
        self._elementalversion = value
    
    @property
    def gaiaVersion(self):
        """ Get gaiaVersion value.

          Notes:
              GaiaVersion is the version of Gaia used by the server.

              
        """
        return self._gaiaversion

    @gaiaVersion.setter
    def gaiaVersion(self, value):
        """ Set gaiaVersion value.

          Notes:
              GaiaVersion is the version of Gaia used by the server.

              
        """
        self._gaiaversion = value
    
    @property
    def manipulateVersion(self):
        """ Get manipulateVersion value.

          Notes:
              ManipulateVersion is the version of Manipulate used by the server.

              
        """
        return self._manipulateversion

    @manipulateVersion.setter
    def manipulateVersion(self, value):
        """ Set manipulateVersion value.

          Notes:
              ManipulateVersion is the version of Manipulate used by the server.

              
        """
        self._manipulateversion = value
    
    @property
    def midgardURL(self):
        """ Get midgardURL value.

          Notes:
              MidgardURL contains the url to use to obtain a token.

              
        """
        return self._midgardurl

    @midgardURL.setter
    def midgardURL(self, value):
        """ Set midgardURL value.

          Notes:
              MidgardURL contains the url to use to obtain a token.

              
        """
        self._midgardurl = value
    
    @property
    def squallVersion(self):
        """ Get squallVersion value.

          Notes:
              SquallVersion is the version of server.

              
        """
        return self._squallversion

    @squallVersion.setter
    def squallVersion(self, value):
        """ Set squallVersion value.

          Notes:
              SquallVersion is the version of server.

              
        """
        self._squallversion = value
    
    @property
    def status(self):
        """ Get status value.

          Notes:
              Status is the overall health status of the server.

              
        """
        return self._status

    @status.setter
    def status(self, value):
        """ Set status value.

          Notes:
              Status is the overall health status of the server.

              
        """
        self._status = value
    
    def validate(self):
        """ Validate valides the current information stored into the structure.
        """
        errors = []

        err = validate_string_in_list("status", self.status, ["Degraded", "Failure", "Ok"], true)

        if err:
            errors.append(err)

        if len(errors) > 0:
            return errors

        return None

    # systeminfoIdentity represents the Identity of the object
systeminfoIdentity = {"name": "systeminfo", "category": "systeminfos", "constructor": SystemInfo}