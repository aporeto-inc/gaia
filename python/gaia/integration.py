# -*- coding: utf-8 -*-

from pyelemental import RESTObject

class Integration(RESTObject):
    """ Represents a Integration in the 

        Notes:
            Integration defines the all the configuration needed to integrate Squall with any 3rd party servers
    """

    def __init__(self, **kwargs):
        """ Initializes a Integration instance

          Notes:
              You can specify all parameters while calling this methods.
              A special argument named `data` will enable you to load the
              object from a Python dictionary

          Examples:
              >>> integration = Integration(id=u'xxxx-xxx-xxx-xxx', name=u'Integration')
              >>> integration = Integration(data=my_dict)
        """

        super(Integration, self).__init__()

        # Read/Write Attributes
        
        self._id = None
        self._annotation = None
        self._associatedtags = None
        self._authtype = None
        self._createdat = None
        self._endpoint = None
        self._namespace = None
        self._normalizedtags = None
        self._parentid = None
        self._parenttype = None
        self._password = None
        self._protected = None
        self._status = None
        self._type = None
        self._updatedat = None
        self._username = None
        
        self.expose_attribute(local_name="ID", remote_name="ID")
        self.expose_attribute(local_name="annotation", remote_name="annotation")
        self.expose_attribute(local_name="associatedTags", remote_name="associatedTags")
        self.expose_attribute(local_name="authType", remote_name="authType")
        self.expose_attribute(local_name="createdAt", remote_name="createdAt")
        self.expose_attribute(local_name="endpoint", remote_name="endpoint")
        self.expose_attribute(local_name="namespace", remote_name="namespace")
        self.expose_attribute(local_name="normalizedTags", remote_name="normalizedTags")
        self.expose_attribute(local_name="parentID", remote_name="parentID")
        self.expose_attribute(local_name="parentType", remote_name="parentType")
        self.expose_attribute(local_name="password", remote_name="password")
        self.expose_attribute(local_name="protected", remote_name="protected")
        self.expose_attribute(local_name="status", remote_name="status")
        self.expose_attribute(local_name="type", remote_name="type")
        self.expose_attribute(local_name="updatedAt", remote_name="updatedAt")
        self.expose_attribute(local_name="userName", remote_name="userName")

        self._compute_args(**kwargs)

    def __str__(self):
        return '<%s:%s>' % (self.identity()["name"], self.identifier)

    def identifier(self):
        """ Identifier returns the value of the object's unique identifier.
        """
        
        return self.ID
        

    def setIdentifier(self, ID):
        """ SetIdentifier sets the value of the object's unique identifier.
        """
        
        self.ID = ID
        

    def identity(self):
        """ Identity returns the Identity of the object.
        """
        return integrationIdentity

    # Properties
    @property
    def ID(self):
        """ Get ID value.

          Notes:
              ID is the identifier of the object.

              
        """
        return self._id

    @ID.setter
    def ID(self, value):
        """ Set ID value.

          Notes:
              ID is the identifier of the object.

              
        """
        self._id = value
    
    @property
    def annotation(self):
        """ Get annotation value.

          Notes:
              Annotation stores additional information about an entity

              
        """
        return self._annotation

    @annotation.setter
    def annotation(self, value):
        """ Set annotation value.

          Notes:
              Annotation stores additional information about an entity

              
        """
        self._annotation = value
    
    @property
    def associatedTags(self):
        """ Get associatedTags value.

          Notes:
              AssociatedTags are the list of tags attached to an entity

              
        """
        return self._associatedtags

    @associatedTags.setter
    def associatedTags(self, value):
        """ Set associatedTags value.

          Notes:
              AssociatedTags are the list of tags attached to an entity

              
        """
        self._associatedtags = value
    
    @property
    def authType(self):
        """ Get authType value.

          Notes:
              AuthType refers to the type of HTTP authentication used to query endpoints

              
        """
        return self._authtype

    @authType.setter
    def authType(self, value):
        """ Set authType value.

          Notes:
              AuthType refers to the type of HTTP authentication used to query endpoints

              
        """
        self._authtype = value
    
    @property
    def createdAt(self):
        """ Get createdAt value.

          Notes:
              CreatedAt is the time at which an entity was created

              
        """
        return self._createdat

    @createdAt.setter
    def createdAt(self, value):
        """ Set createdAt value.

          Notes:
              CreatedAt is the time at which an entity was created

              
        """
        self._createdat = value
    
    @property
    def endpoint(self):
        """ Get endpoint value.

          Notes:
              Endpoint is the API end point of the service

              
        """
        return self._endpoint

    @endpoint.setter
    def endpoint(self, value):
        """ Set endpoint value.

          Notes:
              Endpoint is the API end point of the service

              
        """
        self._endpoint = value
    
    @property
    def namespace(self):
        """ Get namespace value.

          Notes:
              Namespace tag attached to an entity

              
        """
        return self._namespace

    @namespace.setter
    def namespace(self, value):
        """ Set namespace value.

          Notes:
              Namespace tag attached to an entity

              
        """
        self._namespace = value
    
    @property
    def normalizedTags(self):
        """ Get normalizedTags value.

          Notes:
              NormalizedTags contains the list of normalized tags of the entities

              
        """
        return self._normalizedtags

    @normalizedTags.setter
    def normalizedTags(self, value):
        """ Set normalizedTags value.

          Notes:
              NormalizedTags contains the list of normalized tags of the entities

              
        """
        self._normalizedtags = value
    
    @property
    def parentID(self):
        """ Get parentID value.

          Notes:
              ParentID is the ID of the parent, if any,

              
        """
        return self._parentid

    @parentID.setter
    def parentID(self, value):
        """ Set parentID value.

          Notes:
              ParentID is the ID of the parent, if any,

              
        """
        self._parentid = value
    
    @property
    def parentType(self):
        """ Get parentType value.

          Notes:
              ParentType is the type of the parent, if any. It will be set to the parent's Identity.Name.

              
        """
        return self._parenttype

    @parentType.setter
    def parentType(self, value):
        """ Set parentType value.

          Notes:
              ParentType is the type of the parent, if any. It will be set to the parent's Identity.Name.

              
        """
        self._parenttype = value
    
    @property
    def password(self):
        """ Get password value.

          Notes:
              Password is the password of the user to be used in the HTTP Authorization header

              
        """
        return self._password

    @password.setter
    def password(self, value):
        """ Set password value.

          Notes:
              Password is the password of the user to be used in the HTTP Authorization header

              
        """
        self._password = value
    
    @property
    def protected(self):
        """ Get protected value.

          Notes:
              Protected defines if the object is protected.

              
        """
        return self._protected

    @protected.setter
    def protected(self, value):
        """ Set protected value.

          Notes:
              Protected defines if the object is protected.

              
        """
        self._protected = value
    
    @property
    def status(self):
        """ Get status value.

          Notes:
              Status of an entity

              
        """
        return self._status

    @status.setter
    def status(self, value):
        """ Set status value.

          Notes:
              Status of an entity

              
        """
        self._status = value
    
    @property
    def type(self):
        """ Get type value.

          Notes:
              Type refers to type of the server

              
        """
        return self._type

    @type.setter
    def type(self, value):
        """ Set type value.

          Notes:
              Type refers to type of the server

              
        """
        self._type = value
    
    @property
    def updatedAt(self):
        """ Get updatedAt value.

          Notes:
              UpdatedAt is the time at which an entity was updated.

              
        """
        return self._updatedat

    @updatedAt.setter
    def updatedAt(self, value):
        """ Set updatedAt value.

          Notes:
              UpdatedAt is the time at which an entity was updated.

              
        """
        self._updatedat = value
    
    @property
    def userName(self):
        """ Get userName value.

          Notes:
              Username refers to the username to be used in the HTTP Authorization header

              
        """
        return self._username

    @userName.setter
    def userName(self, value):
        """ Set userName value.

          Notes:
              Username refers to the username to be used in the HTTP Authorization header

              
        """
        self._username = value
    

    # integrationIdentity represents the Identity of the object
integrationIdentity = {"name": "integration", "category": "integrations", "constructor": Integration}