# -*- coding: utf-8 -*-

from pyelemental import RESTObject

class APICheck(RESTObject):
    """ Represents a APICheck in the 

        Notes:
            None
    """

    def __init__(self, **kwargs):
        """ Initializes a APICheck instance

          Notes:
              You can specify all parameters while calling this methods.
              A special argument named `data` will enable you to load the
              object from a Python dictionary

          Examples:
              >>> apicheck = APICheck(id=u'xxxx-xxx-xxx-xxx', name=u'APICheck')
              >>> apicheck = APICheck(data=my_dict)
        """

        super(APICheck, self).__init__()

        # Read/Write Attributes
        
        self._authorized = None
        self._namespace = None
        self._operation = None
        self._targetidentities = None
        self._token = None
        
        self.expose_attribute(local_name="authorized", remote_name="authorized")
        self.expose_attribute(local_name="namespace", remote_name="namespace")
        self.expose_attribute(local_name="operation", remote_name="operation")
        self.expose_attribute(local_name="targetIdentities", remote_name="targetIdentities")
        self.expose_attribute(local_name="token", remote_name="token")

        self._compute_args(**kwargs)

    def __str__(self):
        return '<%s:%s>' % (self.identity()["name"], self.identifier)

    def identifier(self):
        """ Identifier returns the value of the object's unique identifier.
        """
        return ""
        

    def setIdentifier(self, ID):
        """ SetIdentifier sets the value of the object's unique identifier.
        """
        pass
        

    def identity(self):
        """ Identity returns the Identity of the object.
        """
        return apicheckIdentity

    # Properties
    @property
    def authorized(self):
        """ Get authorized value.

          Notes:
              Authorized contains the results of the check.

              
        """
        return self._authorized

    @authorized.setter
    def authorized(self, value):
        """ Set authorized value.

          Notes:
              Authorized contains the results of the check.

              
        """
        self._authorized = value
    
    @property
    def namespace(self):
        """ Get namespace value.

          Notes:
              Namespace is the namespace to use to check the api authentication.

              
        """
        return self._namespace

    @namespace.setter
    def namespace(self, value):
        """ Set namespace value.

          Notes:
              Namespace is the namespace to use to check the api authentication.

              
        """
        self._namespace = value
    
    @property
    def operation(self):
        """ Get operation value.

          Notes:
              Operation is the operation you want to check.

              
        """
        return self._operation

    @operation.setter
    def operation(self, value):
        """ Set operation value.

          Notes:
              Operation is the operation you want to check.

              
        """
        self._operation = value
    
    @property
    def targetIdentities(self):
        """ Get targetIdentities value.

          Notes:
              TargetIdentities contains the list of identities you want to check the authorization.

              
        """
        return self._targetidentities

    @targetIdentities.setter
    def targetIdentities(self, value):
        """ Set targetIdentities value.

          Notes:
              TargetIdentities contains the list of identities you want to check the authorization.

              
        """
        self._targetidentities = value
    
    @property
    def token(self):
        """ Get token value.

          Notes:
              Token is the token to use to check api authentication

              
        """
        return self._token

    @token.setter
    def token(self, value):
        """ Set token value.

          Notes:
              Token is the token to use to check api authentication

              
        """
        self._token = value
    

    # apicheckIdentity represents the Identity of the object
apicheckIdentity = {"name": "apicheck", "category": "apichecks", "constructor": APICheck}