# -*- coding: utf-8 -*-

from pyelemental import RESTObject
from pyelemental import validate_string_in_list, validate_float_in_list, validate_int_in_list, validate_required_int, validate_required_float, validate_required_string, validate_required_time, validate_maximum_float, validate_minimum_float, validate_maximum_int, validate_minimum_int, validate_maximum_length, validate_minimum_length, validate_pattern


class ProcessingUnit(RESTObject):
    """ Represents a ProcessingUnit in the 

        Notes:
            ProcessingUnits is the container that gets instantiated on the server.
    """

    def __init__(self, **kwargs):
        """ Initializes a ProcessingUnit instance

          Notes:
              You can specify all parameters while calling this methods.
              A special argument named `data` will enable you to load the
              object from a Python dictionary

          Examples:
              >>> processingunit = ProcessingUnit(id=u'xxxx-xxx-xxx-xxx', name=u'ProcessingUnit')
              >>> processingunit = ProcessingUnit(data=my_dict)
        """

        super(ProcessingUnit, self).__init__()

        # Read/Write Attributes
        
        self._id = None
        self._annotation = None
        self._associatedtags = None
        self._createdat = None
        self._deleted = None
        self._description = None
        self._lastsynctime = None
        self._metadata = None
        self._name = None
        self._namespace = None
        self._nativecontextid = None
        self._normalizedtags = None
        self._operationalstatus = None
        self._parentid = None
        self._parenttype = None
        self._serverid = None
        self._status = None
        self._type = None
        self._updatedat = None
        self._vulnerabilities = None
        
        self.expose_attribute(local_name="ID", remote_name="ID")
        self.expose_attribute(local_name="annotation", remote_name="annotation")
        self.expose_attribute(local_name="associatedTags", remote_name="associatedTags")
        self.expose_attribute(local_name="createdAt", remote_name="createdAt")
        self.expose_attribute(local_name="deleted", remote_name="deleted")
        self.expose_attribute(local_name="description", remote_name="description")
        self.expose_attribute(local_name="lastSyncTime", remote_name="lastSyncTime")
        self.expose_attribute(local_name="metadata", remote_name="metadata")
        self.expose_attribute(local_name="name", remote_name="name")
        self.expose_attribute(local_name="namespace", remote_name="namespace")
        self.expose_attribute(local_name="nativeContextID", remote_name="nativeContextID")
        self.expose_attribute(local_name="normalizedTags", remote_name="normalizedTags")
        self.expose_attribute(local_name="operationalStatus", remote_name="operationalStatus")
        self.expose_attribute(local_name="parentID", remote_name="parentID")
        self.expose_attribute(local_name="parentType", remote_name="parentType")
        self.expose_attribute(local_name="serverID", remote_name="serverID")
        self.expose_attribute(local_name="status", remote_name="status")
        self.expose_attribute(local_name="type", remote_name="type")
        self.expose_attribute(local_name="updatedAt", remote_name="updatedAt")
        self.expose_attribute(local_name="vulnerabilities", remote_name="vulnerabilities")

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
        return processingunitIdentity

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
    def description(self):
        """ Get description value.

          Notes:
              Description is the description of the object.

              
        """
        return self._description

    @description.setter
    def description(self, value):
        """ Set description value.

          Notes:
              Description is the description of the object.

              
        """
        self._description = value
    
    @property
    def lastSyncTime(self):
        """ Get lastSyncTime value.

          Notes:
              LastSyncTime is the time when the policy was last resolved

              
        """
        return self._lastsynctime

    @lastSyncTime.setter
    def lastSyncTime(self, value):
        """ Set lastSyncTime value.

          Notes:
              LastSyncTime is the time when the policy was last resolved

              
        """
        self._lastsynctime = value
    
    @property
    def metadata(self):
        """ Get metadata value.

          Notes:
              Metadata are list of tags associated to the processing unit

              
        """
        return self._metadata

    @metadata.setter
    def metadata(self, value):
        """ Set metadata value.

          Notes:
              Metadata are list of tags associated to the processing unit

              
        """
        self._metadata = value
    
    @property
    def name(self):
        """ Get name value.

          Notes:
              Name is the name of the entity

              
        """
        return self._name

    @name.setter
    def name(self, value):
        """ Set name value.

          Notes:
              Name is the name of the entity

              
        """
        self._name = value
    
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
    def nativeContextID(self):
        """ Get nativeContextID value.

          Notes:
              NativeContextID is the Docker UUID or service PID

              
        """
        return self._nativecontextid

    @nativeContextID.setter
    def nativeContextID(self, value):
        """ Set nativeContextID value.

          Notes:
              NativeContextID is the Docker UUID or service PID

              
        """
        self._nativecontextid = value
    
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
    def operationalStatus(self):
        """ Get operationalStatus value.

          Notes:
              OperationalStatus of the processing unit

              
        """
        return self._operationalstatus

    @operationalStatus.setter
    def operationalStatus(self, value):
        """ Set operationalStatus value.

          Notes:
              OperationalStatus of the processing unit

              
        """
        self._operationalstatus = value
    
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
    def serverID(self):
        """ Get serverID value.

          Notes:
              serverID is the ID of the server associated with the processing unit

              
        """
        return self._serverid

    @serverID.setter
    def serverID(self, value):
        """ Set serverID value.

          Notes:
              serverID is the ID of the server associated with the processing unit

              
        """
        self._serverid = value
    
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
              Type of the container ecosystem

              
        """
        return self._type

    @type.setter
    def type(self, value):
        """ Set type value.

          Notes:
              Type of the container ecosystem

              
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
    def vulnerabilities(self):
        """ Get vulnerabilities value.

          Notes:
              Vulnerabilities contains the list of vulnerabilities of the processing unit.

              
        """
        return self._vulnerabilities

    @vulnerabilities.setter
    def vulnerabilities(self, value):
        """ Set vulnerabilities value.

          Notes:
              Vulnerabilities contains the list of vulnerabilities of the processing unit.

              
        """
        self._vulnerabilities = value
    
    def validate(self):
        """ Validate valides the current information stored into the structure.
        """
        errors = []

        err = validate_required_string("name", self.name)

        if err:
            errors.append(err)

        err = validate_string_in_list("operationalStatus", self.operationalStatus, ["Initialized", "Paused", "Running", "Stopped", "Terminated"], false)

        if err:
            errors.append(err)

        err = validate_string_in_list("type", self.type, ["Docker", "LinuxService", "RKT"], false)

        if err:
            errors.append(err)

        if len(errors) > 0:
            return errors

        return None

    # processingunitIdentity represents the Identity of the object
processingunitIdentity = {"name": "processingunit", "category": "processingunits", "constructor": ProcessingUnit}