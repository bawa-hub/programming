// https://leetcode.com/problems/all-nodes-distance-k-in-binary-tree/

class Solution
{
    void markParents(TreeNode *root, unordered_map<TreeNode *, TreeNode *> &parent_track, TreeNode *target)
    {
        queue<TreeNode *> queue;
        queue.push(root);
        while (!queue.empty())
        {
            TreeNode *current = queue.front();
            queue.pop();
            if (current->left)
            {
                parent_track[current->left] = current;
                queue.push(current->left);
            }
            if (current->right)
            {
                parent_track[current->right] = current;
                queue.push(current->right);
            }
        }
    }

public:
    vector<int> distanceK(TreeNode *root, TreeNode *target, int k)
    {
        unordered_map<TreeNode *, TreeNode *> parent_track; // node -> parent
        markParents(root, parent_track, target);

        unordered_map<TreeNode *, bool> visited;
        queue<TreeNode *> queue;
        queue.push(target);
        visited[target] = true;
        int curr_level = 0;
        while (!queue.empty())
        { /*Second BFS to go upto K level from target node and using our hashtable info*/
            int size = queue.size();
            if (curr_level++ == k)
                break;
            for (int i = 0; i < size; i++)
            {
                TreeNode *current = queue.front();
                queue.pop();
                if (current->left && !visited[current->left])
                {
                    queue.push(current->left);
                    visited[current->left] = true;
                }
                if (current->right && !visited[current->right])
                {
                    queue.push(current->right);
                    visited[current->right] = true;
                }
                if (parent_track[current] && !visited[parent_track[current]])
                {
                    queue.push(parent_track[current]);
                    visited[parent_track[current]] = true;
                }
            }
        }
        vector<int> result;
        while (!queue.empty())
        {
            TreeNode *current = queue.front();
            queue.pop();
            result.push_back(current->val);
        }
        return result;
    }
};

// Time Complexity: O(2N + log N ) The time complexity arises from traversing the tree to create the parent hashmap, which involves visiting every node once hence O(N), exploring all nodes at a distance of ‘K’ which will be O(N) in the worst case, and the logarithmic lookup time for the hashmap is O( log N) in the worst scenario as well hence O(N + N + log N) which simplified to O(N).
// Space Complexity: O(N) The space complexity stems from the data structures used, O(N) for the parent hashmap, O(N) for the queue of DFS, and O(N) for the visited hashmap hence overall our space complexity is O(3N) ~ O(N).