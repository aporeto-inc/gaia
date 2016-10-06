# -*- coding: utf-8 -*-

from pyelemental import RESTObject
from pyelemental import validate_string_in_list, validate_float_in_list, validate_int_in_list, validate_required_int, validate_required_float, validate_required_string, validate_required_time, validate_maximum_float, validate_minimum_float, validate_maximum_int, validate_minimum_int, validate_maximum_length, validate_minimum_length, validate_pattern


class Tag(RESTObject):
    """ Represents a Tag in the 

        Notes:
            A Tag represents a tags associated to an object.
    """

    def __init__(self, **kwargs):
        """ Initializes a Tag instance

          Notes:
              You can specify all parameters while calling this methods.
              A special argument named `data` will enable you to load the
              object from a Python dictionary

          Examples:
              >>> tag = Tag(id=u'xxxx-xxx-xxx-xxx', name=u'Tag')
              >>> tag = Tag(data=my_dict)
        """

        super(Tag, self).__init__()

        # Read/Write Attributes
        
        self._id = None
        self._count = None
        self._namespace = None
        self._value = None
        
        self.expose_attribute(local_name="ID", remote_name="ID")
        self.expose_attribute(local_name="count", remote_name="count")
        self.expose_attribute(local_name="namespace", remote_name="namespace")
        self.expose_attribute(local_name="value", remote_name="value")

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
        return tagIdentity

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
    def count(self):
        """ Get count value.

          Notes:
              Count represents the number of time the tag is used.

              
        """
        return self._count

    @count.setter
    def count(self, value):
        """ Set count value.

          Notes:
              Count represents the number of time the tag is used.

              
        """
        self._count = value
    
    @property
    def namespace(self):
        """ Get namespace value.

          Notes:
              Namespace represents the namespace of the counted tag.

              
        """
        return self._namespace

    @namespace.setter
    def namespace(self, value):
        """ Set namespace value.

          Notes:
              Namespace represents the namespace of the counted tag.

              
        """
        self._namespace = value
    
    @property
    def value(self):
        """ Get value value.

          Notes:
              Value represents the value of the tag.

              
        """
        return self._value

    @value.setter
    def value(self, value):
        """ Set value value.

          Notes:
              Value represents the value of the tag.

              
        """
        self._value = value
    
    def validate(self):
        """ Validate valides the current information stored into the structure.
        """
        errors = []

        err = validate_pattern("value", self.value, "^[^\s=~,]+=[^\s=~,]+$")

        if err:
            errors.append(err)

        err = validate_required_string("value", self.value)

        if err:
            errors.append(err)

        if len(errors) > 0:
            return errors

        return None

    # tagIdentity represents the Identity of the object
tagIdentity = {"name": "tag", "category": "tags", "constructor": Tag}