# -*- coding: utf-8 -*-

from pyelemental import RESTObject

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
        
        self._selector = None
        self._subselectors = None
        self._tonality = None
        
        self.expose_attribute(local_name="selector", remote_name="selector")
        self.expose_attribute(local_name="subSelectors", remote_name="subSelectors")
        self.expose_attribute(local_name="tonality", remote_name="tonality")

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
        return dependencymapsubviewIdentity

    # Properties
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
    

    # dependencymapsubviewIdentity represents the Identity of the object
dependencymapsubviewIdentity = {"name": "dependencymapsubview", "category": "dependencymapsubviews", "constructor": DependencyMapSubview}