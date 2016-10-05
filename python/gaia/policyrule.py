# -*- coding: utf-8 -*-

from pyelemental import RESTObject
from pyelemental import validate_string_in_list, validate_float_in_list, validate_int_in_list, validate_required_int, validate_required_float, validate_required_string, validate_required_time, validate_maximum_float, validate_minimum_float, validate_maximum_int, validate_minimum_int, validate_maximum_length, validate_minimum_length, validate_pattern


class PolicyRule(RESTObject):
    """ Represents a PolicyRule in the 

        Notes:
            PolicyRules describes the set of rules applied on a Subject.
    """

    def __init__(self, **kwargs):
        """ Initializes a PolicyRule instance

          Notes:
              You can specify all parameters while calling this methods.
              A special argument named `data` will enable you to load the
              object from a Python dictionary

          Examples:
              >>> policyrule = PolicyRule(id=u'xxxx-xxx-xxx-xxx', name=u'PolicyRule')
              >>> policyrule = PolicyRule(data=my_dict)
        """

        super(PolicyRule, self).__init__()

        # Read/Write Attributes
        
        self._id = None
        self._action = None
        self._files = None
        self._namespaces = None
        self._networks = None
        self._relation = None
        self._syscalls = None
        self._tagclauses = None
        
        self.expose_attribute(local_name="ID", remote_name="ID")
        self.expose_attribute(local_name="action", remote_name="action")
        self.expose_attribute(local_name="files", remote_name="files")
        self.expose_attribute(local_name="namespaces", remote_name="namespaces")
        self.expose_attribute(local_name="networks", remote_name="networks")
        self.expose_attribute(local_name="relation", remote_name="relation")
        self.expose_attribute(local_name="syscalls", remote_name="syscalls")
        self.expose_attribute(local_name="tagclauses", remote_name="tagclauses")

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
        return policyruleIdentity

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
    def files(self):
        """ Get files value.

          Notes:
              Policy target networks 

              
        """
        return self._files

    @files.setter
    def files(self, value):
        """ Set files value.

          Notes:
              Policy target networks 

              
        """
        self._files = value
    
    @property
    def namespaces(self):
        """ Get namespaces value.

          Notes:
              Policy target networks 

              
        """
        return self._namespaces

    @namespaces.setter
    def namespaces(self, value):
        """ Set namespaces value.

          Notes:
              Policy target networks 

              
        """
        self._namespaces = value
    
    @property
    def networks(self):
        """ Get networks value.

          Notes:
              Policy target networks 

              
        """
        return self._networks

    @networks.setter
    def networks(self, value):
        """ Set networks value.

          Notes:
              Policy target networks 

              
        """
        self._networks = value
    
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
    def syscalls(self):
        """ Get syscalls value.

          Notes:
              Policy target networks 

              
        """
        return self._syscalls

    @syscalls.setter
    def syscalls(self, value):
        """ Set syscalls value.

          Notes:
              Policy target networks 

              
        """
        self._syscalls = value
    
    @property
    def tagclauses(self):
        """ Get tagclauses value.

          Notes:
              Policy target tags

              
        """
        return self._tagclauses

    @tagclauses.setter
    def tagclauses(self, value):
        """ Set tagclauses value.

          Notes:
              Policy target tags

              
        """
        self._tagclauses = value
    
    def validate(self):
        """ Validate valides the current information stored into the structure.
        """
        errors = []

        if len(errors) > 0:
            return errors

        return None

    # policyruleIdentity represents the Identity of the object
policyruleIdentity = {"name": "policyrule", "category": "policyrules", "constructor": PolicyRule}