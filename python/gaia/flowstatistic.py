# -*- coding: utf-8 -*-

from pyelemental import RESTObject

class FlowStatistic(RESTObject):
    """ Represents a FlowStatistic in the 

        Notes:
            FlowStatistic retrieves the flows between two nodes in a specified interval of time.
    """

    def __init__(self, **kwargs):
        """ Initializes a FlowStatistic instance

          Notes:
              You can specify all parameters while calling this methods.
              A special argument named `data` will enable you to load the
              object from a Python dictionary

          Examples:
              >>> flowstatistic = FlowStatistic(id=u'xxxx-xxx-xxx-xxx', name=u'FlowStatistic')
              >>> flowstatistic = FlowStatistic(data=my_dict)
        """

        super(FlowStatistic, self).__init__()

        # Read/Write Attributes
        
        self._id = None
        self._datapoints = None
        self._destinationid = None
        self._destinationtags = None
        self._metric = None
        self._mode = None
        self._sourceid = None
        self._sourcetags = None
        self._type = None
        self._useridentifier = None
        
        self.expose_attribute(local_name="ID", remote_name="ID")
        self.expose_attribute(local_name="dataPoints", remote_name="dataPoints")
        self.expose_attribute(local_name="destinationID", remote_name="destinationID")
        self.expose_attribute(local_name="destinationTags", remote_name="destinationTags")
        self.expose_attribute(local_name="metric", remote_name="metric")
        self.expose_attribute(local_name="mode", remote_name="mode")
        self.expose_attribute(local_name="sourceID", remote_name="sourceID")
        self.expose_attribute(local_name="sourceTags", remote_name="sourceTags")
        self.expose_attribute(local_name="type", remote_name="type")
        self.expose_attribute(local_name="userIdentifier", remote_name="userIdentifier")

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
        return flowstatisticIdentity

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
    def dataPoints(self):
        """ Get dataPoints value.

          Notes:
              DataPoints is a list of time/value pairs that represent the flow events over time.

              
        """
        return self._datapoints

    @dataPoints.setter
    def dataPoints(self, value):
        """ Set dataPoints value.

          Notes:
              DataPoints is a list of time/value pairs that represent the flow events over time.

              
        """
        self._datapoints = value
    
    @property
    def destinationID(self):
        """ Get destinationID value.

          Notes:
              DestinationID is the ID of the destination.

              
        """
        return self._destinationid

    @destinationID.setter
    def destinationID(self, value):
        """ Set destinationID value.

          Notes:
              DestinationID is the ID of the destination.

              
        """
        self._destinationid = value
    
    @property
    def destinationTags(self):
        """ Get destinationTags value.

          Notes:
              DestinationTags contains the tags used to identify destination

              
        """
        return self._destinationtags

    @destinationTags.setter
    def destinationTags(self, value):
        """ Set destinationTags value.

          Notes:
              DestinationTags contains the tags used to identify destination

              
        """
        self._destinationtags = value
    
    @property
    def metric(self):
        """ Get metric value.

          Notes:
              Metric is the kind of metric the statistic represents.

              
        """
        return self._metric

    @metric.setter
    def metric(self, value):
        """ Set metric value.

          Notes:
              Metric is the kind of metric the statistic represents.

              
        """
        self._metric = value
    
    @property
    def mode(self):
        """ Get mode value.

          Notes:
              Mode defines if the metric is for accepted or rejected flows.

              
        """
        return self._mode

    @mode.setter
    def mode(self, value):
        """ Set mode value.

          Notes:
              Mode defines if the metric is for accepted or rejected flows.

              
        """
        self._mode = value
    
    @property
    def sourceID(self):
        """ Get sourceID value.

          Notes:
              SourceID is the source of the stats.

              
        """
        return self._sourceid

    @sourceID.setter
    def sourceID(self, value):
        """ Set sourceID value.

          Notes:
              SourceID is the source of the stats.

              
        """
        self._sourceid = value
    
    @property
    def sourceTags(self):
        """ Get sourceTags value.

          Notes:
              SourceTags contains the tags used to identify the source. 

              
        """
        return self._sourcetags

    @sourceTags.setter
    def sourceTags(self, value):
        """ Set sourceTags value.

          Notes:
              SourceTags contains the tags used to identify the source. 

              
        """
        self._sourcetags = value
    
    @property
    def type(self):
        """ Get type value.

          Notes:
              Type is the type of represenation

              
        """
        return self._type

    @type.setter
    def type(self, value):
        """ Set type value.

          Notes:
              Type is the type of represenation

              
        """
        self._type = value
    
    @property
    def userIdentifier(self):
        """ Get userIdentifier value.

          Notes:
              UserIdentifier can be set by the user as a query parameter. It will be returned in the FlowStatistic object.

              
        """
        return self._useridentifier

    @userIdentifier.setter
    def userIdentifier(self, value):
        """ Set userIdentifier value.

          Notes:
              UserIdentifier can be set by the user as a query parameter. It will be returned in the FlowStatistic object.

              
        """
        self._useridentifier = value
    

    # flowstatisticIdentity represents the Identity of the object
flowstatisticIdentity = {"name": "flowstatistic", "category": "flowstatistics", "constructor": FlowStatistic}