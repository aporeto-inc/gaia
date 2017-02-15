# -*- coding: utf-8 -*-

from pyelemental import RESTObject

class APIAuthorizationPolicy(RESTObject):
    """ Represents a APIAuthorizationPolicy in the 

        Notes:
            None
    """

    def __init__(self, **kwargs):
        """ Initializes a APIAuthorizationPolicy instance

          Notes:
              You can specify all parameters while calling this methods.
              A special argument named `data` will enable you to load the
              object from a Python dictionary

          Examples:
              >>> apiauthorizationpolicy = APIAuthorizationPolicy(id=u'xxxx-xxx-xxx-xxx', name=u'APIAuthorizationPolicy')
              >>> apiauthorizationpolicy = APIAuthorizationPolicy(data=my_dict)
        """

        super(APIAuthorizationPolicy, self).__init__()

        # Read/Write Attributes
        
        self._id = None
        self._allowsdelete = None
        self._allowsget = None
        self._allowshead = None
        self._allowspatch = None
        self._allowspost = None
        self._allowsput = None
        self._annotation = None
        self._associatedtags = None
        self._authorizedidentities = None
        self._authorizednamespace = None
        self._createdat = None
        self._description = None
        self._name = None
        self._namespace = None
        self._normalizedtags = None
        self._parentid = None
        self._parenttype = None
        self._propagate = None
        self._status = None
        self._subject = None
        self._updatedat = None
        
        self.expose_attribute(local_name="ID", remote_name="ID")
        self.expose_attribute(local_name="allowsDelete", remote_name="allowsDelete")
        self.expose_attribute(local_name="allowsGet", remote_name="allowsGet")
        self.expose_attribute(local_name="allowsHead", remote_name="allowsHead")
        self.expose_attribute(local_name="allowsPatch", remote_name="allowsPatch")
        self.expose_attribute(local_name="allowsPost", remote_name="allowsPost")
        self.expose_attribute(local_name="allowsPut", remote_name="allowsPut")
        self.expose_attribute(local_name="annotation", remote_name="annotation")
        self.expose_attribute(local_name="associatedTags", remote_name="associatedTags")
        self.expose_attribute(local_name="authorizedIdentities", remote_name="authorizedIdentities")
        self.expose_attribute(local_name="authorizedNamespace", remote_name="authorizedNamespace")
        self.expose_attribute(local_name="createdAt", remote_name="createdAt")
        self.expose_attribute(local_name="description", remote_name="description")
        self.expose_attribute(local_name="name", remote_name="name")
        self.expose_attribute(local_name="namespace", remote_name="namespace")
        self.expose_attribute(local_name="normalizedTags", remote_name="normalizedTags")
        self.expose_attribute(local_name="parentID", remote_name="parentID")
        self.expose_attribute(local_name="parentType", remote_name="parentType")
        self.expose_attribute(local_name="propagate", remote_name="propagate")
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
        return apiauthorizationpolicyIdentity

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
    def allowsDelete(self):
        """ Get allowsDelete value.

          Notes:
              AllowsDelete defines if DELETE request is authorized.

              
        """
        return self._allowsdelete

    @allowsDelete.setter
    def allowsDelete(self, value):
        """ Set allowsDelete value.

          Notes:
              AllowsDelete defines if DELETE request is authorized.

              
        """
        self._allowsdelete = value
    
    @property
    def allowsGet(self):
        """ Get allowsGet value.

          Notes:
              AllowsGet defines if GET request is authorized.

              
        """
        return self._allowsget

    @allowsGet.setter
    def allowsGet(self, value):
        """ Set allowsGet value.

          Notes:
              AllowsGet defines if GET request is authorized.

              
        """
        self._allowsget = value
    
    @property
    def allowsHead(self):
        """ Get allowsHead value.

          Notes:
              AllowsHead defines if HEAD request is authorized.

              
        """
        return self._allowshead

    @allowsHead.setter
    def allowsHead(self, value):
        """ Set allowsHead value.

          Notes:
              AllowsHead defines if HEAD request is authorized.

              
        """
        self._allowshead = value
    
    @property
    def allowsPatch(self):
        """ Get allowsPatch value.

          Notes:
              AllowsPatch defines if PATCH request is authorized.

              
        """
        return self._allowspatch

    @allowsPatch.setter
    def allowsPatch(self, value):
        """ Set allowsPatch value.

          Notes:
              AllowsPatch defines if PATCH request is authorized.

              
        """
        self._allowspatch = value
    
    @property
    def allowsPost(self):
        """ Get allowsPost value.

          Notes:
              AllowsPost defines if POST request is authorized.

              
        """
        return self._allowspost

    @allowsPost.setter
    def allowsPost(self, value):
        """ Set allowsPost value.

          Notes:
              AllowsPost defines if POST request is authorized.

              
        """
        self._allowspost = value
    
    @property
    def allowsPut(self):
        """ Get allowsPut value.

          Notes:
              AllowsPut defines if PUT request is authorized.

              
        """
        return self._allowsput

    @allowsPut.setter
    def allowsPut(self, value):
        """ Set allowsPut value.

          Notes:
              AllowsPut defines if PUT request is authorized.

              
        """
        self._allowsput = value
    
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
    def authorizedIdentities(self):
        """ Get authorizedIdentities value.

          Notes:
              AuthorizedIdentities defines the list of api identities the policy applies to. 

              
        """
        return self._authorizedidentities

    @authorizedIdentities.setter
    def authorizedIdentities(self, value):
        """ Set authorizedIdentities value.

          Notes:
              AuthorizedIdentities defines the list of api identities the policy applies to. 

              
        """
        self._authorizedidentities = value
    
    @property
    def authorizedNamespace(self):
        """ Get authorizedNamespace value.

          Notes:
              AuthorizedNamespace defines on what namespace the policy applies.

              
        """
        return self._authorizednamespace

    @authorizedNamespace.setter
    def authorizedNamespace(self, value):
        """ Set authorizedNamespace value.

          Notes:
              AuthorizedNamespace defines on what namespace the policy applies.

              
        """
        self._authorizednamespace = value
    
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
    def propagate(self):
        """ Get propagate value.

          Notes:
              Propagate defines if the policy should propagate.

              
        """
        return self._propagate

    @propagate.setter
    def propagate(self, value):
        """ Set propagate value.

          Notes:
              Propagate defines if the policy should propagate.

              
        """
        self._propagate = value
    
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
              Subject is the subject.

              
        """
        return self._subject

    @subject.setter
    def subject(self, value):
        """ Set subject value.

          Notes:
              Subject is the subject.

              
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
    

    # apiauthorizationpolicyIdentity represents the Identity of the object
apiauthorizationpolicyIdentity = {"name": "apiauthorizationpolicy", "category": "apiauthorizationpolicies", "constructor": APIAuthorizationPolicy}