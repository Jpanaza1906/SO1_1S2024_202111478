import React, { useEffect, useRef } from 'react';
import Chart from 'chart.js/auto';

const PieChart = ({ data, labels, colors,title}) => {
  let chartRef = useRef(null);

  useEffect(() => {
    let myChart = null;
    if (chartRef.current && myChart) {
      myChart.destroy(); // Si el gráfico ya existe, destrúyelo
    }

    const ctx = chartRef.current.getContext('2d');
    myChart = new Chart(ctx, {
      type: 'pie',
      data: {
        labels: labels,
        datasets: [{
          label: title,
          data: data,
          backgroundColor: colors,
          borderColor: colors,
          hoverBackgroundColor: colors, //['#269fbd', '#87f4e8'],
            borderWidth: 1
        }]
      },
      options: {
        responsive: true,
        maintainAspectRatio: false,
        plugins: {
          legend: {
            position: 'top',
            animatiom: false
          },
          title: {
            display: true,
            text: title
          }
        },
        animation: {
          duration: 0
        }
      }
    });
    return () => {
      if (myChart) {
        myChart.destroy(); // Asegúrate de destruir el gráfico al desmontar el componente
      }
    };
  }, [data, labels, colors,title]);

  return (
    <div>
      <canvas ref={chartRef}></canvas>
    </div>
  );
};

export default PieChart;
