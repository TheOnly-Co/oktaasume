apt-get update && \
apt-get install -y --allow-change-held-packages kubeadm=1.19.0-00

apt-get update && \
    apt-get install -y --allow-change-held-packages kubelet=1.19.0-00 kubectl=1.19.0-00

sudo kubeadm upgrade apply v1.19.0



kubectl drain node01 --ignore-daemonsets


ssh node01

    apt-get update && \
apt-get install -y --allow-change-held-packages kubeadm=1.19.0-00

apt-get update && \
apt-get install -y --allow-change-held-packages kubelet=1.19.0-00 kubectl=1.19.0-00

kubeadm upgrade node

logout

sudo systemctl daemon-reload
sudo systemctl restart kubelet

kubectl uncordon node01
