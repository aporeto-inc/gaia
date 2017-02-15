# -*- coding: utf-8 -*-

from pyelemental import RESTObject

class RenderedPolicy(RESTObject):
    """ Represents a RenderedPolicy in the 

        Notes:
            RenderedPolicies attached to the given set of Subjects.
    """

    def __init__(self, **kwargs):
        """ Initializes a RenderedPolicy instance

          Notes:
              You can specify all parameters while calling this methods.
              A special argument named `data` will enable you to load the
              object from a Python dictionary

          Examples:
              >>> renderedpolicy = RenderedPolicy(id=u'xxxx-xxx-xxx-xxx', name=u'RenderedPolicy')
              >>> renderedpolicy = RenderedPolicy(data=my_dict)
        """

        super(RenderedPolicy, self).__init__()

        # Read/Write Attributes
        
        self._id = None
        self._egresspolicies = None
        self._ingresspolicies = None
        self._processingunitid = None
        self._profile = None
        
        self.expose_attribute(local_name="ID", remote_name="ID")
        self.expose_attribute(local_name="egressPolicies", remote_name="egressPolicies")
        self.expose_attribute(local_name="ingressPolicies", remote_name="ingressPolicies")
        self.expose_attribute(local_name="processingUnitID", remote_name="processingUnitID")
        self.expose_attribute(local_name="profile", remote_name="profile")

        self._compute_args(**kwargs)

    def __str__(self):
        return '<%s:%s>' % (self.identity()["name"], self.identifier)

    def identifier(self):
        """ Identifier returns the value of the object's unique identifier.
        """
        
        return self.ProcessingUnitID
        

    def setIdentifier(self, ID):
        """ SetIdentifier sets the value of the object's unique identifier.
        """
        
        self.ProcessingUnitID = ID
        

    def identity(self):
        """ Identity returns the Identity of the object.
        """
        return renderedpolicyIdentity

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
    def egressPolicies(self):
        """ Get egressPolicies value.

          Notes:
              EgressPolicies lists all the egress policies attached to ProcessingUnit

              
        """
        return self._egresspolicies

    @egressPolicies.setter
    def egressPolicies(self, value):
        """ Set egressPolicies value.

          Notes:
              EgressPolicies lists all the egress policies attached to ProcessingUnit

              
        """
        self._egresspolicies = value
    
    @property
    def ingressPolicies(self):
        """ Get ingressPolicies value.

          Notes:
              IngressPolicies lists all the ingress policies attached to ProcessingUnit

              
        """
        return self._ingresspolicies

    @ingressPolicies.setter
    def ingressPolicies(self, value):
        """ Set ingressPolicies value.

          Notes:
              IngressPolicies lists all the ingress policies attached to ProcessingUnit

              
        """
        self._ingresspolicies = value
    
    @property
    def processingUnitID(self):
        """ Get processingUnitID value.

          Notes:
              Identifier of the ProcessingUnit

              
        """
        return self._processingunitid

    @processingUnitID.setter
    def processingUnitID(self, value):
        """ Set processingUnitID value.

          Notes:
              Identifier of the ProcessingUnit

              
        """
        self._processingunitid = value
    
    @property
    def profile(self):
        """ Get profile value.

          Notes:
              Profile is the trust profile of the processing unit that should be used during all communications.

              
        """
        return self._profile

    @profile.setter
    def profile(self, value):
        """ Set profile value.

          Notes:
              Profile is the trust profile of the processing unit that should be used during all communications.

              
        """
        self._profile = value
    

    # renderedpolicyIdentity represents the Identity of the object
renderedpolicyIdentity = {"name": "renderedpolicy", "category": "renderedpolicies", "constructor": RenderedPolicy}