# -*- coding: utf-8 -*-

from pyelemental import RESTObject
from pyelemental import validate_string_in_list, validate_float_in_list, validate_int_in_list, validate_required_int, validate_required_float, validate_required_string, validate_required_time, validate_maximum_float, validate_minimum_float, validate_maximum_int, validate_minimum_int, validate_maximum_length, validate_minimum_length, validate_pattern


class MapEdge(RESTObject):
    """ Represents a MapEdge in the 

        Notes:
            MapEdge describes a dependency between two resources.
    """

    def __init__(self, **kwargs):
        """ Initializes a MapEdge instance

          Notes:
              You can specify all parameters while calling this methods.
              A special argument named `data` will enable you to load the
              object from a Python dictionary

          Examples:
              >>> mapedge = MapEdge(id=u'xxxx-xxx-xxx-xxx', name=u'MapEdge')
              >>> mapedge = MapEdge(data=my_dict)
        """

        super(MapEdge, self).__init__()

        # Read/Write Attributes
        
        self._id = None
        self._destinationid = None
        self._labels = None
        self._name = None
        self._sourceid = None
        
        self.expose_attribute(local_name="ID", remote_name="ID")
        self.expose_attribute(local_name="destinationID", remote_name="destinationID")
        self.expose_attribute(local_name="labels", remote_name="labels")
        self.expose_attribute(local_name="name", remote_name="name")
        self.expose_attribute(local_name="sourceID", remote_name="sourceID")

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
        return mapedgeIdentity

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
    def destinationID(self):
        """ Get destinationID value.

          Notes:
              ID of the destination resource

              
        """
        return self._destinationid

    @destinationID.setter
    def destinationID(self, value):
        """ Set destinationID value.

          Notes:
              ID of the destination resource

              
        """
        self._destinationid = value
    
    @property
    def labels(self):
        """ Get labels value.

          Notes:
              Labels provide grouping parameters

              
        """
        return self._labels

    @labels.setter
    def labels(self, value):
        """ Set labels value.

          Notes:
              Labels provide grouping parameters

              
        """
        self._labels = value
    
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
    def sourceID(self):
        """ Get sourceID value.

          Notes:
              ID of the source resource 

              
        """
        return self._sourceid

    @sourceID.setter
    def sourceID(self, value):
        """ Set sourceID value.

          Notes:
              ID of the source resource 

              
        """
        self._sourceid = value
    
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

    # mapedgeIdentity represents the Identity of the object
mapedgeIdentity = {"name": "mapedge", "category": "mapedges", "constructor": MapEdge}