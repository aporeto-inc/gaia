# -*- coding: utf-8 -*-

from pyelemental import RESTObject

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
        self._certificateauthority = None
        self._elementalversion = None
        self._gaiaversion = None
        self._googleclientid = None
        self._kairosdburl = None
        self._manipulateversion = None
        self._midgardurl = None
        self._pubsubservice = None
        self._squallversion = None
        self._status = None
        self._vinceurl = None
        self._zackurl = None
        
        self.expose_attribute(local_name="APIVersion", remote_name="APIVersion")
        self.expose_attribute(local_name="bahamutVersion", remote_name="bahamutVersion")
        self.expose_attribute(local_name="certificateAuthority", remote_name="certificateAuthority")
        self.expose_attribute(local_name="elementalVersion", remote_name="elementalVersion")
        self.expose_attribute(local_name="gaiaVersion", remote_name="gaiaVersion")
        self.expose_attribute(local_name="googleClientID", remote_name="googleClientID")
        self.expose_attribute(local_name="kairosDBURL", remote_name="kairosDBURL")
        self.expose_attribute(local_name="manipulateVersion", remote_name="manipulateVersion")
        self.expose_attribute(local_name="midgardURL", remote_name="midgardURL")
        self.expose_attribute(local_name="pubSubService", remote_name="pubSubService")
        self.expose_attribute(local_name="squallVersion", remote_name="squallVersion")
        self.expose_attribute(local_name="status", remote_name="status")
        self.expose_attribute(local_name="vinceURL", remote_name="vinceURL")
        self.expose_attribute(local_name="zackURL", remote_name="zackURL")

        self._compute_args(**kwargs)

    def __str__(self):
        return '<%s:%s>' % (self.identity()["name"], self.identifier)

    def identifier(self):
        """ Identifier returns the value of the object's unique identifier.
        """
        return ""
        

    def setIdentifier(self, ID):
        """ SetIdentifier sets the value of the object's unique identifier.
        """
        pass
        

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
    def certificateAuthority(self):
        """ Get certificateAuthority value.

          Notes:
              CertificateAuthority contains the main certificate authority,

              
        """
        return self._certificateauthority

    @certificateAuthority.setter
    def certificateAuthority(self, value):
        """ Set certificateAuthority value.

          Notes:
              CertificateAuthority contains the main certificate authority,

              
        """
        self._certificateauthority = value
    
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
    def googleClientID(self):
        """ Get googleClientID value.

          Notes:
              GoogleClientID is the Google oauth client ID to use to get a valid token from Google for Midgard.

              
        """
        return self._googleclientid

    @googleClientID.setter
    def googleClientID(self, value):
        """ Set googleClientID value.

          Notes:
              GoogleClientID is the Google oauth client ID to use to get a valid token from Google for Midgard.

              
        """
        self._googleclientid = value
    
    @property
    def kairosDBURL(self):
        """ Get kairosDBURL value.

          Notes:
              KairosDBURL contains the URL of the kairos db server.

              
        """
        return self._kairosdburl

    @kairosDBURL.setter
    def kairosDBURL(self, value):
        """ Set kairosDBURL value.

          Notes:
              KairosDBURL contains the URL of the kairos db server.

              
        """
        self._kairosdburl = value
    
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
    def pubSubService(self):
        """ Get pubSubService value.

          Notes:
              PubsubService provides the end-point for the pubsub server.

              
        """
        return self._pubsubservice

    @pubSubService.setter
    def pubSubService(self, value):
        """ Set pubSubService value.

          Notes:
              PubsubService provides the end-point for the pubsub server.

              
        """
        self._pubsubservice = value
    
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
    
    @property
    def vinceURL(self):
        """ Get vinceURL value.

          Notes:
              VinceURL contains the URL for the Vince server.

              
        """
        return self._vinceurl

    @vinceURL.setter
    def vinceURL(self, value):
        """ Set vinceURL value.

          Notes:
              VinceURL contains the URL for the Vince server.

              
        """
        self._vinceurl = value
    
    @property
    def zackURL(self):
        """ Get zackURL value.

          Notes:
              zackURL contains the URL of the Zack server.

              
        """
        return self._zackurl

    @zackURL.setter
    def zackURL(self, value):
        """ Set zackURL value.

          Notes:
              zackURL contains the URL of the Zack server.

              
        """
        self._zackurl = value
    

    # systeminfoIdentity represents the Identity of the object
systeminfoIdentity = {"name": "systeminfo", "category": "systeminfos", "constructor": SystemInfo}