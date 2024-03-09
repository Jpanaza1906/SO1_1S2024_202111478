import React, { useEffect, useRef } from "react";
import { DataSet, Network } from "vis-network/standalone/umd/vis-network.min";

const State = ({ action }) => {
    const containerRef = useRef(null);
    let networkRef = useRef(null);
    const nodes = useRef(null);
    const edges = useRef(null);

    useEffect(() => {
        // Define los nodos 
        nodes.current = new DataSet([]);

        // Define las aristas
        edges.current = new DataSet([]);

        // Crea una instancia de la red
        const container = containerRef.current;
        const data = {
            nodes: nodes.current,
            edges: edges.current
        };
        const options = {
            layout: {
                hierarchical: {
                    direction: 'LR'
                }
            },
            edges:{
                smooth:{
                    type: 'curvedCCW'
                }
            }
        };
        networkRef.current = new Network(container, data, options);

        // Limpia el diagrama cuando se desmonta el componente
        return () => {
            if (networkRef.current) {
                networkRef.current.destroy();
            }
        };
    }, []);

    useEffect(() => {


        // Si la accion es start se borra el diagrama y se agregan los nodos
        if (action === 'start') {
            nodes.current.clear();
            nodes.current.add([
                { id: 1, label: 'New', color: 'gray' },
                { id: 2, label: 'Ready', color: 'gray' },
                { id: 3, label: 'Running', color: 'green' }
            ]);
            // Se agregan las aristas
            edges.current.clear();
            edges.current.add({ from: 1, to: 2, arrows: { to: { enabled: true, type: 'arrow' } } });
            edges.current.add({ from: 2, to: 3, arrows: { to: { enabled: true, type: 'arrow' } } });
        }

        //actualiza el color del nodo 'Running'
        if (action === 'stop' || action === 'resume') {
            const runningNode = nodes.current.get(3);
            if (runningNode) {
                runningNode.color = action === 'stop' ? 'red' : 'green';
                nodes.current.update(runningNode);
            }
        }

        // añade una arista de running a ready cuando se recibe stop
        if (action === 'stop') {
            // Agrega una nueva arista de 'Running' a 'Ready'
            edges.current.add({ from: 3, to: 2, arrows: { to: { enabled: true, type: 'arrow' } } });
            

            // Actualiza el color del nodo 'Ready' a verde
            const readyNode = nodes.current.get(2);
            if (readyNode) {
                readyNode.color = 'green';
                nodes.current.update(readyNode);
            }
        }

        // elimina la arista de running a ready cuando se recibe resume
        if (action === 'resume') {
            // Busca la arista existente de 'Running' a 'Ready'
            const existingEdge = edges.current.get().find(edge => edge.from === 3 && edge.to === 2);
            if (existingEdge) {
                // Elimina la arista existente de 'Running' a 'Ready'
                edges.current.remove(existingEdge.id);

                // Agrega una nueva arista de 'Ready' a 'Running'
                //edges.current.add({ from: 2, to: 3, arrows: { to: { enabled: true, type: 'arrow' } } });
            }

            // Actualiza el color del nodo 'Ready' a gris
            const readyNode = nodes.current.get(2);
            if (readyNode) {
                readyNode.color = 'gray';
                nodes.current.update(readyNode);
            }
        }

        // elimina la arista de running a ready  cuando se recibe kill
        if (action === 'kill') {
            let targetNode;
            // Verifica el color del nodo 'Ready'
            const readyNode = nodes.current.get(2);
            if (readyNode && readyNode.color === 'green') {
                targetNode = 2; // Si el nodo 'Ready' está en verde, este será el nodo objetivo
            }
            // Si el nodo 'Ready' no está en verde, verifica el color del nodo 'Running'
            else {
                const runningNode = nodes.current.get(3);
                if (runningNode && runningNode.color === 'green') {
                    targetNode = 3; // Si el nodo 'Running' está en verde, este será el nodo objetivo
                }
            }

            // Todos los nodos se ponen en azul
            nodes.current.forEach(node => {
                node.color = 'gray';
                nodes.current.update(node);
            });

            // Se agrega el nodo 'Terminated'
            nodes.current.add({ id: 4, label: 'Terminated', color: 'orange' });

            // Si se identificó un nodo objetivo, se crea la arista de ese nodo al nodo 'Terminated'
            if (targetNode) {
                edges.current.add({ from: targetNode, to: 4, arrows: { to: { enabled: true, type: 'arrow' } } });
            }
        }

    }, [action]);

    return (
        <div>
            <div ref={containerRef} style={{ width: '100%', height: '550px' }} ></div>
        </div>
    );
}

export default State;