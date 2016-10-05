# -*- coding: utf-8 -*-

from pyelemental import RESTObject
from pyelemental import validate_string_in_list, validate_float_in_list, validate_int_in_list, validate_required_int, validate_required_float, validate_required_string, validate_required_time, validate_maximum_float, validate_minimum_float, validate_maximum_int, validate_minimum_int, validate_maximum_length, validate_minimum_length, validate_pattern


class ComputedPolicy(RESTObject):
    """ Represents a ComputedPolicy in the 

        Notes:
            Policies computed by the system from a dependency map view
    """

    def __init__(self, **kwargs):
        """ Initializes a ComputedPolicy instance

          Notes:
              You can specify all parameters while calling this methods.
              A special argument named `data` will enable you to load the
              object from a Python dictionary

          Examples:
              >>> computedpolicy = ComputedPolicy(id=u'xxxx-xxx-xxx-xxx', name=u'ComputedPolicy')
              >>> computedpolicy = ComputedPolicy(data=my_dict)
        """

        super(ComputedPolicy, self).__init__()

        # Read/Write Attributes
        
        self._id = None
        self._networkaccesspolicies = None
        
        self.expose_attribute(local_name="ID", remote_name="ID")
        self.expose_attribute(local_name="networkAccessPolicies", remote_name="networkAccessPolicies")

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
        return computedpolicyIdentity

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
    def networkAccessPolicies(self):
        """ Get networkAccessPolicies value.

          Notes:
              Array of netowrk access policies computed

              
        """
        return self._networkaccesspolicies

    @networkAccessPolicies.setter
    def networkAccessPolicies(self, value):
        """ Set networkAccessPolicies value.

          Notes:
              Array of netowrk access policies computed

              
        """
        self._networkaccesspolicies = value
    
    def validate(self):
        """ Validate valides the current information stored into the structure.
        """
        errors = []

        if len(errors) > 0:
            return errors

        return None

    # computedpolicyIdentity represents the Identity of the object
computedpolicyIdentity = {"name": "computedpolicy", "category": "computedpolicies", "constructor": ComputedPolicy}