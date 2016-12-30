# -*- coding: utf-8 -*-

from pyelemental import RESTObject
from pyelemental import validate_string_in_list, validate_float_in_list, validate_int_in_list, validate_required_int, validate_required_float, validate_required_string, validate_required_time, validate_maximum_float, validate_minimum_float, validate_maximum_int, validate_minimum_int, validate_maximum_length, validate_minimum_length, validate_pattern


class Policy(RESTObject):
    """ Represents a Policy in the 

        Notes:
            Policy that describes the set of rules to be applied when an event happens in the system. For example, policy can be defined to authorize an API call.
    """

    def __init__(self, **kwargs):
        """ Initializes a Policy instance

          Notes:
              You can specify all parameters while calling this methods.
              A special argument named `data` will enable you to load the
              object from a Python dictionary

          Examples:
              >>> policy = Policy(id=u'xxxx-xxx-xxx-xxx', name=u'Policy')
              >>> policy = Policy(data=my_dict)
        """

        super(Policy, self).__init__()

        # Read/Write Attributes
        
        self._id = None
        self._action = None
        self._allobjecttags = None
        self._allsubjecttags = None
        self._annotation = None
        self._associatedtags = None
        self._createdat = None
        self._deleted = None
        self._description = None
        self._name = None
        self._namespace = None
        self._normalizedtags = None
        self._object = None
        self._parentid = None
        self._parenttype = None
        self._propagate = None
        self._relation = None
        self._status = None
        self._subject = None
        self._type = None
        self._updatedat = None
        
        self.expose_attribute(local_name="ID", remote_name="ID")
        self.expose_attribute(local_name="action", remote_name="action")
        self.expose_attribute(local_name="allObjectTags", remote_name="allObjectTags")
        self.expose_attribute(local_name="allSubjectTags", remote_name="allSubjectTags")
        self.expose_attribute(local_name="annotation", remote_name="annotation")
        self.expose_attribute(local_name="associatedTags", remote_name="associatedTags")
        self.expose_attribute(local_name="createdAt", remote_name="createdAt")
        self.expose_attribute(local_name="deleted", remote_name="deleted")
        self.expose_attribute(local_name="description", remote_name="description")
        self.expose_attribute(local_name="name", remote_name="name")
        self.expose_attribute(local_name="namespace", remote_name="namespace")
        self.expose_attribute(local_name="normalizedTags", remote_name="normalizedTags")
        self.expose_attribute(local_name="object", remote_name="object")
        self.expose_attribute(local_name="parentID", remote_name="parentID")
        self.expose_attribute(local_name="parentType", remote_name="parentType")
        self.expose_attribute(local_name="propagate", remote_name="propagate")
        self.expose_attribute(local_name="relation", remote_name="relation")
        self.expose_attribute(local_name="status", remote_name="status")
        self.expose_attribute(local_name="subject", remote_name="subject")
        self.expose_attribute(local_name="type", remote_name="type")
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
        return policyIdentity

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
    def action(self):
        """ Get action value.

          Notes:
              Action defines set of actions that must be enforced when a dependency is met.

              
        """
        return self._action

    @action.setter
    def action(self, value):
        """ Set action value.

          Notes:
              Action defines set of actions that must be enforced when a dependency is met.

              
        """
        self._action = value
    
    @property
    def allObjectTags(self):
        """ Get allObjectTags value.

          Notes:
              This is a set of all object tags for matching in the DB

              
        """
        return self._allobjecttags

    @allObjectTags.setter
    def allObjectTags(self, value):
        """ Set allObjectTags value.

          Notes:
              This is a set of all object tags for matching in the DB

              
        """
        self._allobjecttags = value
    
    @property
    def allSubjectTags(self):
        """ Get allSubjectTags value.

          Notes:
              This is a set of all subject tags for matching in the DB

              
        """
        return self._allsubjecttags

    @allSubjectTags.setter
    def allSubjectTags(self, value):
        """ Set allSubjectTags value.

          Notes:
              This is a set of all subject tags for matching in the DB

              
        """
        self._allsubjecttags = value
    
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
              Object represents set of entities that another entity depends on. As subjects, objects are identified as logical operations on tags when a policy is defined.

              
        """
        return self._object

    @object.setter
    def object(self, value):
        """ Set object value.

          Notes:
              Object represents set of entities that another entity depends on. As subjects, objects are identified as logical operations on tags when a policy is defined.

              
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
    def propagate(self):
        """ Get propagate value.

          Notes:
              Propagate will propagate the policy to all of its children.

              
        """
        return self._propagate

    @propagate.setter
    def propagate(self, value):
        """ Set propagate value.

          Notes:
              Propagate will propagate the policy to all of its children.

              
        """
        self._propagate = value
    
    @property
    def relation(self):
        """ Get relation value.

          Notes:
              Relation describes the required operation to be performed between subjects and objects

              
        """
        return self._relation

    @relation.setter
    def relation(self, value):
        """ Set relation value.

          Notes:
              Relation describes the required operation to be performed between subjects and objects

              
        """
        self._relation = value
    
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
              Subject represent sets of entities that will have a dependency other entities. Subjects are defined as logical operations on tags. Logical operations can includes AND/OR

              
        """
        return self._subject

    @subject.setter
    def subject(self, value):
        """ Set subject value.

          Notes:
              Subject represent sets of entities that will have a dependency other entities. Subjects are defined as logical operations on tags. Logical operations can includes AND/OR

              
        """
        self._subject = value
    
    @property
    def type(self):
        """ Get type value.

          Notes:
              Type of the policy

              
        """
        return self._type

    @type.setter
    def type(self, value):
        """ Set type value.

          Notes:
              Type of the policy

              
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
    
    def validate(self):
        """ Validate valides the current information stored into the structure.
        """
        errors = []

        err = validate_required_string("name", self.name)

        if err:
            errors.append(err)

        if len(errors) > 0:
            return errors

        return None

    # policyIdentity represents the Identity of the object
policyIdentity = {"name": "policy", "category": "policies", "constructor": Policy}