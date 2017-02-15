# -*- coding: utf-8 -*-

from pyelemental import RESTObject

class FileAccessPolicy(RESTObject):
    """ Represents a FileAccessPolicy in the 

        Notes:
            None
    """

    def __init__(self, **kwargs):
        """ Initializes a FileAccessPolicy instance

          Notes:
              You can specify all parameters while calling this methods.
              A special argument named `data` will enable you to load the
              object from a Python dictionary

          Examples:
              >>> fileaccesspolicy = FileAccessPolicy(id=u'xxxx-xxx-xxx-xxx', name=u'FileAccessPolicy')
              >>> fileaccesspolicy = FileAccessPolicy(data=my_dict)
        """

        super(FileAccessPolicy, self).__init__()

        # Read/Write Attributes
        
        self._id = None
        self._propagate = None
        self._allowsexecute = None
        self._allowsread = None
        self._allowswrite = None
        self._annotation = None
        self._associatedtags = None
        self._createdat = None
        self._description = None
        self._encryptionenabled = None
        self._logsenabled = None
        self._name = None
        self._namespace = None
        self._normalizedtags = None
        self._object = None
        self._parentid = None
        self._parenttype = None
        self._status = None
        self._subject = None
        self._updatedat = None
        
        self.expose_attribute(local_name="ID", remote_name="ID")
        self.expose_attribute(local_name="Propagate", remote_name="Propagate")
        self.expose_attribute(local_name="allowsExecute", remote_name="allowsExecute")
        self.expose_attribute(local_name="allowsRead", remote_name="allowsRead")
        self.expose_attribute(local_name="allowsWrite", remote_name="allowsWrite")
        self.expose_attribute(local_name="annotation", remote_name="annotation")
        self.expose_attribute(local_name="associatedTags", remote_name="associatedTags")
        self.expose_attribute(local_name="createdAt", remote_name="createdAt")
        self.expose_attribute(local_name="description", remote_name="description")
        self.expose_attribute(local_name="encryptionEnabled", remote_name="encryptionEnabled")
        self.expose_attribute(local_name="logsEnabled", remote_name="logsEnabled")
        self.expose_attribute(local_name="name", remote_name="name")
        self.expose_attribute(local_name="namespace", remote_name="namespace")
        self.expose_attribute(local_name="normalizedTags", remote_name="normalizedTags")
        self.expose_attribute(local_name="object", remote_name="object")
        self.expose_attribute(local_name="parentID", remote_name="parentID")
        self.expose_attribute(local_name="parentType", remote_name="parentType")
        self.expose_attribute(local_name="status", remote_name="status")
        self.expose_attribute(local_name="subject", remote_name="subject")
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
        return fileaccesspolicyIdentity

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
    def Propagate(self):
        """ Get Propagate value.

          Notes:
              Propagate indicates that the policy must be propagated to the children namespaces.

              
        """
        return self._propagate

    @Propagate.setter
    def Propagate(self, value):
        """ Set Propagate value.

          Notes:
              Propagate indicates that the policy must be propagated to the children namespaces.

              
        """
        self._propagate = value
    
    @property
    def allowsExecute(self):
        """ Get allowsExecute value.

          Notes:
              AllowsExecute allows to execute the files.

              
        """
        return self._allowsexecute

    @allowsExecute.setter
    def allowsExecute(self, value):
        """ Set allowsExecute value.

          Notes:
              AllowsExecute allows to execute the files.

              
        """
        self._allowsexecute = value
    
    @property
    def allowsRead(self):
        """ Get allowsRead value.

          Notes:
              AllowsRead allows to read the files.

              
        """
        return self._allowsread

    @allowsRead.setter
    def allowsRead(self, value):
        """ Set allowsRead value.

          Notes:
              AllowsRead allows to read the files.

              
        """
        self._allowsread = value
    
    @property
    def allowsWrite(self):
        """ Get allowsWrite value.

          Notes:
              AllowsWrite allows to write the files.

              
        """
        return self._allowswrite

    @allowsWrite.setter
    def allowsWrite(self, value):
        """ Set allowsWrite value.

          Notes:
              AllowsWrite allows to write the files.

              
        """
        self._allowswrite = value
    
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
    def encryptionEnabled(self):
        """ Get encryptionEnabled value.

          Notes:
              EncryptionEnabled will enable the automatic encryption

              
        """
        return self._encryptionenabled

    @encryptionEnabled.setter
    def encryptionEnabled(self, value):
        """ Set encryptionEnabled value.

          Notes:
              EncryptionEnabled will enable the automatic encryption

              
        """
        self._encryptionenabled = value
    
    @property
    def logsEnabled(self):
        """ Get logsEnabled value.

          Notes:
              LogsEnabled will enable logging when this policy is used.

              
        """
        return self._logsenabled

    @logsEnabled.setter
    def logsEnabled(self, value):
        """ Set logsEnabled value.

          Notes:
              LogsEnabled will enable logging when this policy is used.

              
        """
        self._logsenabled = value
    
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
    def object(self):
        """ Get object value.

          Notes:
              Object is the object of the policy.

              
        """
        return self._object

    @object.setter
    def object(self, value):
        """ Set object value.

          Notes:
              Object is the object of the policy.

              
        """
        self._object = value
    
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
    def subject(self):
        """ Get subject value.

          Notes:
              Subject is the subject of the policy

              
        """
        return self._subject

    @subject.setter
    def subject(self, value):
        """ Set subject value.

          Notes:
              Subject is the subject of the policy

              
        """
        self._subject = value
    
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
    

    # fileaccesspolicyIdentity represents the Identity of the object
fileaccesspolicyIdentity = {"name": "fileaccesspolicy", "category": "fileaccesspolicies", "constructor": FileAccessPolicy}