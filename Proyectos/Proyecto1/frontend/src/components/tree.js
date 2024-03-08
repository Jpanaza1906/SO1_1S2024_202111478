import React, { useEffect, useRef } from "react";
import { DataSet, Network } from "vis-network/standalone/umd/vis-network.min";

const Tree = ({ processData }) => {
    const containerRef = useRef(null);
    let networkRef = useRef(null);

    useEffect(() => {
        if (processData && containerRef.current) {
            const nodes = new DataSet();
            const edges = new DataSet();

            nodes.add({ id: processData.pid, label: `${processData.name}\npid = ${processData.pid}` });

            if (processData.child) {

                processData.child.forEach(child => {
                    nodes.add({ id: child.pid, label: `${child.name}\npid = ${child.pid}` });
                    edges.add({ from: child.pidPadre, to: child.pid });
                });
            }

            const data = {
                nodes: nodes,
                edges: edges
            };

            const options = {
                layout: {
                    hierarchical: {
                        direction: "UD",
                        sortMethod: "directed",
                    },
                },
                edges: {
                    font: { align: "top" }
                }
            };

            networkRef.current = new Network(containerRef.current, data, options);
        }

        return () => {
            if (networkRef.current) {
                networkRef.current.destroy();
            }
        };
    }, [processData]);

    return (
        <div>
            <div ref={containerRef} style={{ width: '100%', height: '550px' }} ></div>
        </div>
    );

};

export default Tree;