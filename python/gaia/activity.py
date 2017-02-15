# -*- coding: utf-8 -*-

from pyelemental import RESTObject

class Activity(RESTObject):
    """ Represents a Activity in the 

        Notes:
            None
    """

    def __init__(self, **kwargs):
        """ Initializes a Activity instance

          Notes:
              You can specify all parameters while calling this methods.
              A special argument named `data` will enable you to load the
              object from a Python dictionary

          Examples:
              >>> activity = Activity(id=u'xxxx-xxx-xxx-xxx', name=u'Activity')
              >>> activity = Activity(data=my_dict)
        """

        super(Activity, self).__init__()

        # Read/Write Attributes
        
        self._id = None
        self._claims = None
        self._data = None
        self._date = None
        self._error = None
        self._message = None
        self._namespace = None
        self._operation = None
        self._targetidentity = None
        
        self.expose_attribute(local_name="ID", remote_name="ID")
        self.expose_attribute(local_name="claims", remote_name="claims")
        self.expose_attribute(local_name="data", remote_name="data")
        self.expose_attribute(local_name="date", remote_name="date")
        self.expose_attribute(local_name="error", remote_name="error")
        self.expose_attribute(local_name="message", remote_name="message")
        self.expose_attribute(local_name="namespace", remote_name="namespace")
        self.expose_attribute(local_name="operation", remote_name="operation")
        self.expose_attribute(local_name="targetIdentity", remote_name="targetIdentity")

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
        return activityIdentity

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
    def claims(self):
        """ Get claims value.

          Notes:
              Claims of the user who performed the operation.

              
        """
        return self._claims

    @claims.setter
    def claims(self, value):
        """ Set claims value.

          Notes:
              Claims of the user who performed the operation.

              
        """
        self._claims = value
    
    @property
    def data(self):
        """ Get data value.

          Notes:
              Data of the notification.

              
        """
        return self._data

    @data.setter
    def data(self, value):
        """ Set data value.

          Notes:
              Data of the notification.

              
        """
        self._data = value
    
    @property
    def date(self):
        """ Get date value.

          Notes:
              Date of the notification.

              
        """
        return self._date

    @date.setter
    def date(self, value):
        """ Set date value.

          Notes:
              Date of the notification.

              
        """
        self._date = value
    
    @property
    def error(self):
        """ Get error value.

          Notes:
              Error contains the eventual error.

              
        """
        return self._error

    @error.setter
    def error(self, value):
        """ Set error value.

          Notes:
              Error contains the eventual error.

              
        """
        self._error = value
    
    @property
    def message(self):
        """ Get message value.

          Notes:
              Message of the notification.

              
        """
        return self._message

    @message.setter
    def message(self, value):
        """ Set message value.

          Notes:
              Message of the notification.

              
        """
        self._message = value
    
    @property
    def namespace(self):
        """ Get namespace value.

          Notes:
              Namespace of the notification.

              
        """
        return self._namespace

    @namespace.setter
    def namespace(self, value):
        """ Set namespace value.

          Notes:
              Namespace of the notification.

              
        """
        self._namespace = value
    
    @property
    def operation(self):
        """ Get operation value.

          Notes:
              Operation describe what kind of operation the notification represents.

              
        """
        return self._operation

    @operation.setter
    def operation(self, value):
        """ Set operation value.

          Notes:
              Operation describe what kind of operation the notification represents.

              
        """
        self._operation = value
    
    @property
    def targetIdentity(self):
        """ Get targetIdentity value.

          Notes:
              TargetIdentity is the Identity of the related object.

              
        """
        return self._targetidentity

    @targetIdentity.setter
    def targetIdentity(self, value):
        """ Set targetIdentity value.

          Notes:
              TargetIdentity is the Identity of the related object.

              
        """
        self._targetidentity = value
    

    # activityIdentity represents the Identity of the object
activityIdentity = {"name": "activity", "category": "activities", "constructor": Activity}