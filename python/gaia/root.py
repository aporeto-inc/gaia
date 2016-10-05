# -*- coding: utf-8 -*-

from pyelemental import RESTObject
from pyelemental import validate_string_in_list, validate_float_in_list, validate_int_in_list, validate_required_int, validate_required_float, validate_required_string, validate_required_time, validate_maximum_float, validate_minimum_float, validate_maximum_int, validate_minimum_int, validate_maximum_length, validate_minimum_length, validate_pattern


class Root(RESTObject):
    """ Represents a Root in the 

        Notes:
            None
    """

    def __init__(self, **kwargs):
        """ Initializes a Root instance

          Notes:
              You can specify all parameters while calling this methods.
              A special argument named `data` will enable you to load the
              object from a Python dictionary

          Examples:
              >>> root = Root(id=u'xxxx-xxx-xxx-xxx', name=u'Root')
              >>> root = Root(data=my_dict)
        """

        super(Root, self).__init__()

        # Read/Write Attributes
        
        self._id = None
        
        self.expose_attribute(local_name="ID", remote_name="ID")

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
        return rootIdentity

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
    
    def validate(self):
        """ Validate valides the current information stored into the structure.
        """
        errors = []

        if len(errors) > 0:
            return errors

        return None

    def APIKey(self):
        """ APIKey returns a the API Key
        """
        return self.Token

    def setAPIKey(self, key):
        """ SetAPIKey sets a the API Key
        """
        self.Token = key

    # rootIdentity represents the Identity of the object
rootIdentity = {"name": "root", "category": "root", "constructor": Root}