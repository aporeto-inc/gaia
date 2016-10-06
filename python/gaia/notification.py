# -*- coding: utf-8 -*-

from pyelemental import RESTObject
from pyelemental import validate_string_in_list, validate_float_in_list, validate_int_in_list, validate_required_int, validate_required_float, validate_required_string, validate_required_time, validate_maximum_float, validate_minimum_float, validate_maximum_int, validate_minimum_int, validate_maximum_length, validate_minimum_length, validate_pattern


class Notification(RESTObject):
    """ Represents a Notification in the 

        Notes:
            None
    """

    def __init__(self, **kwargs):
        """ Initializes a Notification instance

          Notes:
              You can specify all parameters while calling this methods.
              A special argument named `data` will enable you to load the
              object from a Python dictionary

          Examples:
              >>> notification = Notification(id=u'xxxx-xxx-xxx-xxx', name=u'Notification')
              >>> notification = Notification(data=my_dict)
        """

        super(Notification, self).__init__()

        # Read/Write Attributes
        
        self._id = None
        self._created = None
        self._deleted = None
        self._layerlimit = None
        self._name = None
        self._new = None
        self._nextpage = None
        self._notified = None
        self._old = None
        self._page = None
        
        self.expose_attribute(local_name="ID", remote_name="ID")
        self.expose_attribute(local_name="created", remote_name="created")
        self.expose_attribute(local_name="deleted", remote_name="deleted")
        self.expose_attribute(local_name="layerLimit", remote_name="layerLimit")
        self.expose_attribute(local_name="name", remote_name="name")
        self.expose_attribute(local_name="new", remote_name="new")
        self.expose_attribute(local_name="nextPage", remote_name="nextPage")
        self.expose_attribute(local_name="notified", remote_name="notified")
        self.expose_attribute(local_name="old", remote_name="old")
        self.expose_attribute(local_name="page", remote_name="page")

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
        return notificationIdentity

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
    def created(self):
        """ Get created value.

          Notes:
              Created is the time when then notification was created

              
        """
        return self._created

    @created.setter
    def created(self, value):
        """ Set created value.

          Notes:
              Created is the time when then notification was created

              
        """
        self._created = value
    
    @property
    def deleted(self):
        """ Get deleted value.

          Notes:
              Deleted is the time when the notification was deleted

              
        """
        return self._deleted

    @deleted.setter
    def deleted(self, value):
        """ Set deleted value.

          Notes:
              Deleted is the time when the notification was deleted

              
        """
        self._deleted = value
    
    @property
    def layerLimit(self):
        """ Get layerLimit value.

          Notes:
              LayerLimit is the number of layers returned in notification

              
        """
        return self._layerlimit

    @layerLimit.setter
    def layerLimit(self, value):
        """ Set layerLimit value.

          Notes:
              LayerLimit is the number of layers returned in notification

              
        """
        self._layerlimit = value
    
    @property
    def name(self):
        """ Get name value.

          Notes:
              Name is the name of the notification

              
        """
        return self._name

    @name.setter
    def name(self, value):
        """ Set name value.

          Notes:
              Name is the name of the notification

              
        """
        self._name = value
    
    @property
    def new(self):
        """ Get new value.

          Notes:
              New is the new layers that introduced vulnerability

              
        """
        return self._new

    @new.setter
    def new(self, value):
        """ Set new value.

          Notes:
              New is the new layers that introduced vulnerability

              
        """
        self._new = value
    
    @property
    def nextPage(self):
        """ Get nextPage value.

          Notes:
              NextPage is the next page number

              
        """
        return self._nextpage

    @nextPage.setter
    def nextPage(self, value):
        """ Set nextPage value.

          Notes:
              NextPage is the next page number

              
        """
        self._nextpage = value
    
    @property
    def notified(self):
        """ Get notified value.

          Notes:
              Notified is the time when the notification was sent

              
        """
        return self._notified

    @notified.setter
    def notified(self, value):
        """ Set notified value.

          Notes:
              Notified is the time when the notification was sent

              
        """
        self._notified = value
    
    @property
    def old(self):
        """ Get old value.

          Notes:
              Old is the old layers that introduced vulnerability

              
        """
        return self._old

    @old.setter
    def old(self, value):
        """ Set old value.

          Notes:
              Old is the old layers that introduced vulnerability

              
        """
        self._old = value
    
    @property
    def page(self):
        """ Get page value.

          Notes:
              Page is the page number

              
        """
        return self._page

    @page.setter
    def page(self, value):
        """ Set page value.

          Notes:
              Page is the page number

              
        """
        self._page = value
    
    def validate(self):
        """ Validate valides the current information stored into the structure.
        """
        errors = []

        if len(errors) > 0:
            return errors

        return None

    # notificationIdentity represents the Identity of the object
notificationIdentity = {"name": "notification", "category": "notifications", "constructor": Notification}