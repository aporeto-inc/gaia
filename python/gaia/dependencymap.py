# -*- coding: utf-8 -*-

from pyelemental import RESTObject
from pyelemental import validate_string_in_list, validate_float_in_list, validate_int_in_list, validate_required_int, validate_required_float, validate_required_string, validate_required_time, validate_maximum_float, validate_minimum_float, validate_maximum_int, validate_minimum_int, validate_maximum_length, validate_minimum_length, validate_pattern


class DependencyMap(RESTObject):
    """ Represents a DependencyMap in the 

        Notes:
            dependencymap creates a map of dependencies.
    """

    def __init__(self, **kwargs):
        """ Initializes a DependencyMap instance

          Notes:
              You can specify all parameters while calling this methods.
              A special argument named `data` will enable you to load the
              object from a Python dictionary

          Examples:
              >>> dependencymap = DependencyMap(id=u'xxxx-xxx-xxx-xxx', name=u'DependencyMap')
              >>> dependencymap = DependencyMap(data=my_dict)
        """

        super(DependencyMap, self).__init__()

        # Read/Write Attributes
        
        self._id = None
        self._edges = None
        self._groups = None
        self._nodes = None
        
        self.expose_attribute(local_name="ID", remote_name="ID")
        self.expose_attribute(local_name="edges", remote_name="edges")
        self.expose_attribute(local_name="groups", remote_name="groups")
        self.expose_attribute(local_name="nodes", remote_name="nodes")

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
        return dependencymapIdentity

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
    def edges(self):
        """ Get edges value.

          Notes:
              edges are the edges of the map

              
        """
        return self._edges

    @edges.setter
    def edges(self, value):
        """ Set edges value.

          Notes:
              edges are the edges of the map

              
        """
        self._edges = value
    
    @property
    def groups(self):
        """ Get groups value.

          Notes:
              Groups provide information about the group values

              
        """
        return self._groups

    @groups.setter
    def groups(self, value):
        """ Set groups value.

          Notes:
              Groups provide information about the group values

              
        """
        self._groups = value
    
    @property
    def nodes(self):
        """ Get nodes value.

          Notes:
              nodes refers to the nodes of the map

              
        """
        return self._nodes

    @nodes.setter
    def nodes(self, value):
        """ Set nodes value.

          Notes:
              nodes refers to the nodes of the map

              
        """
        self._nodes = value
    
    def validate(self):
        """ Validate valides the current information stored into the structure.
        """
        errors = []

        if len(errors) > 0:
            return errors

        return None

    # dependencymapIdentity represents the Identity of the object
dependencymapIdentity = {"name": "dependencymap", "category": "dependencymaps", "constructor": DependencyMap}