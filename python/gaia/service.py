# -*- coding: utf-8 -*-

from pyelemental import RESTObject
from pyelemental import validate_string_in_list, validate_float_in_list, validate_int_in_list, validate_required_int, validate_required_float, validate_required_string, validate_required_time, validate_maximum_float, validate_minimum_float, validate_maximum_int, validate_minimum_int, validate_maximum_length, validate_minimum_length, validate_pattern


class Service(RESTObject):
    """ Represents a Service in the 

        Notes:
            None
    """

    def __init__(self, **kwargs):
        """ Initializes a Service instance

          Notes:
              You can specify all parameters while calling this methods.
              A special argument named `data` will enable you to load the
              object from a Python dictionary

          Examples:
              >>> service = Service(id=u'xxxx-xxx-xxx-xxx', name=u'Service')
              >>> service = Service(data=my_dict)
        """

        super(Service, self).__init__()

        # Read/Write Attributes
        
        self._id = None
        self._annotation = None
        self._associatedtags = None
        self._createdat = None
        self._deleted = None
        self._endpoint = None
        self._namespace = None
        self._parentid = None
        self._parenttype = None
        self._port = None
        self._server = None
        self._ssl = None
        self._status = None
        self._updatedat = None
        
        self.expose_attribute(local_name="ID", remote_name="ID")
        self.expose_attribute(local_name="annotation", remote_name="annotation")
        self.expose_attribute(local_name="associatedTags", remote_name="associatedTags")
        self.expose_attribute(local_name="createdAt", remote_name="createdAt")
        self.expose_attribute(local_name="deleted", remote_name="deleted")
        self.expose_attribute(local_name="endpoint", remote_name="endpoint")
        self.expose_attribute(local_name="namespace", remote_name="namespace")
        self.expose_attribute(local_name="parentID", remote_name="parentID")
        self.expose_attribute(local_name="parentType", remote_name="parentType")
        self.expose_attribute(local_name="port", remote_name="port")
        self.expose_attribute(local_name="server", remote_name="server")
        self.expose_attribute(local_name="ssl", remote_name="ssl")
        self.expose_attribute(local_name="status", remote_name="status")
        self.expose_attribute(local_name="updatedAt", remote_name="updatedAt")

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
        return serviceIdentity

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
    def deleted(self):
        """ Get deleted value.

          Notes:
              Deleted marks if the entity has been deleted.

              
        """
        return self._deleted

    @deleted.setter
    def deleted(self, value):
        """ Set deleted value.

          Notes:
              Deleted marks if the entity has been deleted.

              
        """
        self._deleted = value
    
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
    def port(self):
        """ Get port value.

          Notes:
              Port is the port number of the service

              
        """
        return self._port

    @port.setter
    def port(self, value):
        """ Set port value.

          Notes:
              Port is the port number of the service

              
        """
        self._port = value
    
    @property
    def server(self):
        """ Get server value.

          Notes:
              Server is either the DNS name or IP of the server that provides the service

              
        """
        return self._server

    @server.setter
    def server(self, value):
        """ Set server value.

          Notes:
              Server is either the DNS name or IP of the server that provides the service

              
        """
        self._server = value
    
    @property
    def ssl(self):
        """ Get ssl value.

          Notes:
              SSL defines if the service is either secured or unsecured

              
        """
        return self._ssl

    @ssl.setter
    def ssl(self, value):
        """ Set ssl value.

          Notes:
              SSL defines if the service is either secured or unsecured

              
        """
        self._ssl = value
    
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
    
    def validate(self):
        """ Validate valides the current information stored into the structure.
        """
        errors = []

        err = validate_string_in_list("ssl", self.ssl, ["Disabled", "Enabled"], false)

        if err:
            errors.append(err)

        if len(errors) > 0:
            return errors

        return None

    # serviceIdentity represents the Identity of the object
serviceIdentity = {"name": "service", "category": "services", "constructor": Service}