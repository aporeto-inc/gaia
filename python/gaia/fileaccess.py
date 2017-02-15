# -*- coding: utf-8 -*-

from pyelemental import RESTObject

class FileAccess(RESTObject):
    """ Represents a FileAccess in the 

        Notes:
            None
    """

    def __init__(self, **kwargs):
        """ Initializes a FileAccess instance

          Notes:
              You can specify all parameters while calling this methods.
              A special argument named `data` will enable you to load the
              object from a Python dictionary

          Examples:
              >>> fileaccess = FileAccess(id=u'xxxx-xxx-xxx-xxx', name=u'FileAccess')
              >>> fileaccess = FileAccess(data=my_dict)
        """

        super(FileAccess, self).__init__()

        # Read/Write Attributes
        
        self._action = None
        self._count = None
        self._host = None
        self._mode = None
        self._path = None
        self._protocol = None
        
        self.expose_attribute(local_name="action", remote_name="action")
        self.expose_attribute(local_name="count", remote_name="count")
        self.expose_attribute(local_name="host", remote_name="host")
        self.expose_attribute(local_name="mode", remote_name="mode")
        self.expose_attribute(local_name="path", remote_name="path")
        self.expose_attribute(local_name="protocol", remote_name="protocol")

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
        return fileaccessIdentity

    # Properties
    @property
    def action(self):
        """ Get action value.

          Notes:
              Action tells if the access has been allowed or not.

              
        """
        return self._action

    @action.setter
    def action(self, value):
        """ Set action value.

          Notes:
              Action tells if the access has been allowed or not.

              
        """
        self._action = value
    
    @property
    def count(self):
        """ Get count value.

          Notes:
              Count tells how many times the file has been accessed.

              
        """
        return self._count

    @count.setter
    def count(self, value):
        """ Set count value.

          Notes:
              Count tells how many times the file has been accessed.

              
        """
        self._count = value
    
    @property
    def host(self):
        """ Get host value.

          Notes:
              Host is the host that served the accessed file.

              
        """
        return self._host

    @host.setter
    def host(self, value):
        """ Set host value.

          Notes:
              Host is the host that served the accessed file.

              
        """
        self._host = value
    
    @property
    def mode(self):
        """ Get mode value.

          Notes:
              Mode is the mode of the accessed file.

              
        """
        return self._mode

    @mode.setter
    def mode(self, value):
        """ Set mode value.

          Notes:
              Mode is the mode of the accessed file.

              
        """
        self._mode = value
    
    @property
    def path(self):
        """ Get path value.

          Notes:
              Path is the path of the accessed file.

              
        """
        return self._path

    @path.setter
    def path(self, value):
        """ Set path value.

          Notes:
              Path is the path of the accessed file.

              
        """
        self._path = value
    
    @property
    def protocol(self):
        """ Get protocol value.

          Notes:
              Protocol is the protocol used to access the file.

              
        """
        return self._protocol

    @protocol.setter
    def protocol(self, value):
        """ Set protocol value.

          Notes:
              Protocol is the protocol used to access the file.

              
        """
        self._protocol = value
    

    # fileaccessIdentity represents the Identity of the object
fileaccessIdentity = {"name": "fileaccess", "category": "fileaccesses", "constructor": FileAccess}