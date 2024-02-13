import React, { useState, useEffect, useRef, memo } from 'react';
import Chart from 'chart.js/auto';
import { GetRAMData } from "../wailsjs/go/main/App";
import './App.css';

function App() {
    const [memoryData, setMemoryData] = useState(null);
    const chartContainerRef = useRef(null);
    const chartInstanceRef = useRef(null);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await GetRAMData();
                const parsedData = JSON.parse(response);
                console.log("Data received:", parsedData)
                setMemoryData(parsedData);
            } catch (error) {
                console.error("Error fetching data:", error);
            }
        };

        const createChart = () => {
            console.log("Creating or updating chart...");
            console.log("Memory data:", memoryData);
            console.log("Chart container ref:", chartContainerRef.current);
            if (memoryData && chartContainerRef.current) {
                const ctx = chartContainerRef.current.getContext('2d');
                const chartConfig = {
                    type: 'pie',
                    data: {
                        datasets: [{
                            data: [memoryData.free_memory, memoryData.used_memory],
                            backgroundColor: ['#00c6ff', '#37FF8B'],
                            hoverBackgroundColor: ['#00c6ff', '#37FF8B']
                        }],
                        labels: ['Used Memory', 'Free Memory']
                    }
                };
                if (!chartInstanceRef.current) {
                    console.log("Creating new chart instance...");
                    chartInstanceRef.current = new Chart(ctx, chartConfig);
                } else {
                    console.log("Updating existing chart instance...");
                    chartInstanceRef.current.data.datasets[0].data = [memoryData.used_memory, memoryData.free_memory];
                    chartInstanceRef.current.update();
                }
            }
        };

        const interval = setInterval(() => {
            fetchData();
        }, 5000);

        if (memoryData) {
            createChart();
        }

        return () => clearInterval(interval);
    }, [memoryData]); // Solo se ejecuta una vez al montar el componente

    return (
        <div className="App">
            <header className="App-header">
            <button class="button" data-text="Awesome">
                <span class="actual-text">&nbsp;RAM&nbsp;</span>
                <span aria-hidden="true" class="hover-text">&nbsp;RAM&nbsp;</span>
            </button>
                <div className='Chart'>
                    <canvas ref={chartContainerRef} />
                </div>
                {memoryData && (
                    <h3>
                        {`${(memoryData.used_memory / 1000000000).toFixed(1)} of ${(memoryData.total_memory / 1000000000).toFixed(1)} GB`}
                        {' '}
                        {`(${memoryData.percentage} %)`}
                    </h3>
                )}
            </header>
        </div>
    );
}

export default App;
