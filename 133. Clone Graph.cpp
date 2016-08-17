class Solution {
public:
    map<UndirectedGraphNode *,UndirectedGraphNode *> m;
    UndirectedGraphNode *cloneGraph(UndirectedGraphNode *node) {
        if(!node) return NULL;
        if(m.find(node) == m.end()){
            m[node] = new UndirectedGraphNode(node->label);
            for(UndirectedGraphNode* neighbor : node->neighbors){
                m[node]->neighbors.push_back(cloneGraph(neighbor));
            }
        }
        return m[node];
    }
};