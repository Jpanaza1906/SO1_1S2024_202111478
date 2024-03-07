import React, { useEffect, useRef } from 'react';
import Chart from 'chart.js/auto';
import '../css/linechart.css';


const LineChart = ({ data}) => {
    let chartRef = useRef(null);
    useEffect(() => {
      let myChart = null;
      if (chartRef.current && myChart) {
        myChart.destroy();
      }
  
      const ctx = chartRef.current.getContext('2d');
      myChart = new Chart(ctx, {
        type: 'line',
        data: data,
        options: {
          responsive: true,
          maintainAspectRatio: false,
          animation: {
            duration: 0
          },
          scales: {
            y: {
              beginAtZero: true,
              max: 100
            }
          }
        }
      });
  
      return () => {
        if (myChart) {
          myChart.destroy();
        }
      };
    }, [data]);
  
    return (
      <div className='chartline-container'>
        <canvas ref={chartRef}></canvas>
      </div>
    );
  };

export default LineChart;