# -*- coding: utf-8 -*-

from pyelemental import RESTObject
from pyelemental import validate_string_in_list, validate_float_in_list, validate_int_in_list, validate_required_int, validate_required_float, validate_required_string, validate_required_time, validate_maximum_float, validate_minimum_float, validate_maximum_int, validate_minimum_int, validate_maximum_length, validate_minimum_length, validate_pattern


class DependencyMapSubview(RESTObject):
    """ Represents a DependencyMapSubview in the 

        Notes:
            A DependencyMapSubview is subview of a DependencyMapView
    """

    def __init__(self, **kwargs):
        """ Initializes a DependencyMapSubview instance

          Notes:
              You can specify all parameters while calling this methods.
              A special argument named `data` will enable you to load the
              object from a Python dictionary

          Examples:
              >>> dependencymapsubview = DependencyMapSubview(id=u'xxxx-xxx-xxx-xxx', name=u'DependencyMapSubview')
              >>> dependencymapsubview = DependencyMapSubview(data=my_dict)
        """

        super(DependencyMapSubview, self).__init__()

        # Read/Write Attributes
        
        self._id = None
        self._selector = None
        self._subselectors = None
        self._tonality = None
        
        self.expose_attribute(local_name="ID", remote_name="ID")
        self.expose_attribute(local_name="selector", remote_name="selector")
        self.expose_attribute(local_name="subSelectors", remote_name="subSelectors")
        self.expose_attribute(local_name="tonality", remote_name="tonality")

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
        return dependencymapsubviewIdentity

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
    def selector(self):
        """ Get selector value.

          Notes:
              Selector is the main selector for the DependencyMapSubview.

              
        """
        return self._selector

    @selector.setter
    def selector(self, value):
        """ Set selector value.

          Notes:
              Selector is the main selector for the DependencyMapSubview.

              
        """
        self._selector = value
    
    @property
    def subSelectors(self):
        """ Get subSelectors value.

          Notes:
              SubSelectors are the selector to apply inside the main selector.

              
        """
        return self._subselectors

    @subSelectors.setter
    def subSelectors(self, value):
        """ Set subSelectors value.

          Notes:
              SubSelectors are the selector to apply inside the main selector.

              
        """
        self._subselectors = value
    
    @property
    def tonality(self):
        """ Get tonality value.

          Notes:
              Tonality sets the color tonality to use for the DependencyMapSubView.

              
        """
        return self._tonality

    @tonality.setter
    def tonality(self, value):
        """ Set tonality value.

          Notes:
              Tonality sets the color tonality to use for the DependencyMapSubView.

              
        """
        self._tonality = value
    
    def validate(self):
        """ Validate valides the current information stored into the structure.
        """
        errors = []

        if len(errors) > 0:
            return errors

        return None

    # dependencymapsubviewIdentity represents the Identity of the object
dependencymapsubviewIdentity = {"name": "dependencymapsubview", "category": "dependencymapsubviews", "constructor": DependencyMapSubview}