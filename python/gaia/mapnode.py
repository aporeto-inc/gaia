# -*- coding: utf-8 -*-

from pyelemental import RESTObject
from pyelemental import validate_string_in_list, validate_float_in_list, validate_int_in_list, validate_required_int, validate_required_float, validate_required_string, validate_required_time, validate_maximum_float, validate_minimum_float, validate_maximum_int, validate_minimum_int, validate_maximum_length, validate_minimum_length, validate_pattern


class MapNode(RESTObject):
    """ Represents a MapNode in the 

        Notes:
            MapNode describes a resource for the map.
    """

    def __init__(self, **kwargs):
        """ Initializes a MapNode instance

          Notes:
              You can specify all parameters while calling this methods.
              A special argument named `data` will enable you to load the
              object from a Python dictionary

          Examples:
              >>> mapnode = MapNode(id=u'xxxx-xxx-xxx-xxx', name=u'MapNode')
              >>> mapnode = MapNode(data=my_dict)
        """

        super(MapNode, self).__init__()

        # Read/Write Attributes
        
        self._id = None
        self._groups = None
        self._name = None
        self._type = None
        
        self.expose_attribute(local_name="ID", remote_name="ID")
        self.expose_attribute(local_name="groups", remote_name="groups")
        self.expose_attribute(local_name="name", remote_name="name")
        self.expose_attribute(local_name="type", remote_name="type")

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
        return mapnodeIdentity

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
    def groups(self):
        """ Get groups value.

          Notes:
              Groups for organizing resources

              
        """
        return self._groups

    @groups.setter
    def groups(self, value):
        """ Set groups value.

          Notes:
              Groups for organizing resources

              
        """
        self._groups = value
    
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
    def type(self):
        """ Get type value.

          Notes:
              Type of the resource represented in the map

              
        """
        return self._type

    @type.setter
    def type(self, value):
        """ Set type value.

          Notes:
              Type of the resource represented in the map

              
        """
        self._type = value
    
    def validate(self):
        """ Validate valides the current information stored into the structure.
        """
        errors = []

        err = validate_required_string("name", self.name)

        if err:
            errors.append(err)

        err = validate_string_in_list("type", self.type, ["Container", "Volume"], true)

        if err:
            errors.append(err)

        if len(errors) > 0:
            return errors

        return None

    # mapnodeIdentity represents the Identity of the object
mapnodeIdentity = {"name": "mapnode", "category": "mapnodes", "constructor": MapNode}