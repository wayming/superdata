import React from 'react';
import './App.css';
import '../node_modules/nvd3/build/nv.d3.css';
import NVD3Chart from "react-nvd3";
// import d3 from "d3";
 import { useState } from 'react';
var datum = [{
  key: "sun",
  values: [ 
    { 'vdate': '2015-06-10T00:00:00.000Z', 'value': 1.18665 },
    { 'vdate': '2015-06-11T00:00:00.000Z', 'value': 1.19168 },
    { 'vdate': '2015-06-12T00:00:00.000Z', 'value': 1.18928 },
    { 'vdate': '2015-06-13T00:00:00.000Z', 'value': 1.18928 },
    { 'vdate': '2015-06-14T00:00:00.000Z', 'value': 1.18928 },
    { 'vdate': '2015-06-15T00:00:00.000Z', 'value': 1.18543 },
    { 'vdate': '2015-06-16T00:00:00.000Z', 'value': 1.18601 },
    { 'vdate': '2015-06-17T00:00:00.000Z', 'value': 1.19112 },
    { 'vdate': '2015-06-18T00:00:00.000Z', 'value': 1.18608 },
    { 'vdate': '2015-06-19T00:00:00.000Z', 'value': 1.19101 },
    { 'vdate': '2015-06-20T00:00:00.000Z', 'value': 1.19101 },
    { 'vdate': '2015-06-21T00:00:00.000Z', 'value': 1.19101 },
    { 'vdate': '2015-06-22T00:00:00.000Z', 'value': 1.19509 },
    { 'vdate': '2015-06-23T00:00:00.000Z', 'value': 1.20044 },
  ]
}];

function Chart() {
  const [units, setUnits] = useState(false);
  console.log(datum)
  fetch('http://localhost:8001')
    .then(response => {
      return response.text();
    })
    .then(data => {
      console.log(data)
      setUnits(data);
    });
  return (
    <div className="Chart">
        {
          React.createElement(NVD3Chart, {
            xAxis: {
              tickFormat: function(d){ return Date(d); },
              axisLabel: 'Period'
            },
            yAxis: {
              tickFormat: function(d) {return parseFloat(d).toFixed(6); }
            },
            type:'cumulativeLineChart',
            datum: datum,
            x: function(d) {
              return new Date(d['vdate']);
            },
            y: function(d) {
              return d['value'];
            },
            height: 600
          })
        }
      <h1> { units } } </h1>

    </div>
  );
}

export default Chart;
