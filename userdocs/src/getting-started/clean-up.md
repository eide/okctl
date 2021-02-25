During testing, or maybe when a project has come to an end - there might be a necessity to delete a cluster created with
`okctl`

## Delete a cluster

```bash
# Delete an okctl cluster. Format:
# okctl delete cluster <environment name>
#
# <cluster name>     must equal the environment name used for cluster creation.
#
# Example:
okctl delete cluster test
```

### Delete hosted zone

If you want to remove your cluster's primary hosted zone, use

```bash
okctl delete cluster <environment name> --i-know-what-i-am-doing-delete-hosted-zone-and-records true 
```

Note: When deleting the hosted zone, you are deleting your NS record. The NS record contains a
a TTL setting which determines how long cache resolvers cache your NS record. If the TTL is 15
minutes, you should wait 15 minutes before creating a new cluster - if you're using the same
domain name.

## Delete a cluster manually

The delete operation may have failed for several reasons. In the following sections, we list common causes of problems
and possible remedies.

### Delete all Kubernetes resources

Delete everything in all namespaces except `kube-system`. In `kube-system`, some of the core components, such as 
`AWS load balancer controller` are running. These controllers create AWS resources on your behalf. By removing all 
resources from all namespaces, we allow these services to clean up themselves.

```bash
for each in $(kubectl get ns -o jsonpath="{.items[*].metadata.name}" | grep -v kube-system);
do
  kubectl delete ns $each
done
```

### Remove Fargate Profile

Log in to the AWS console for your account, then go to `Elastic Kubernetes Service > Clusters`. Select your cluster, 
then in the `Configuration` find the `Compute` sub tab. Find the `Fargate Profiles` items and delete the `fp-default` 
profile.

### Remove generated route53 records

Log in to the AWS console for your account, the go to `route 53` and select the relevant hosted zone. The only two
records that should remain so that Cloud Formation is able to delete it is the SOA record and the NS record.

### Delete the cloud formation stacks in reverse order

Log in to the AWS console for your account, then go to `Cloud Formation > Stacks`. If your AWS account contains many 
stacks, filter the list by `okctl` and the cluster name and environment.

Start deleting the stacks in order from _newest_ to _oldest_, if a stack ends up in a state where the delete operation
fails. Take a closer look at the events generated by the delete operation on the stack. Frequently you will find a 
resource that has failed to delete for some reason. Follow this resource to its corresponding page. Maybe a subnet has 
failed to delete because there are active network interfaces on it. By following the resource and trying to delete it 
manually, you will be given additional information. Perhaps you need to disconnect the interfaces from an EC2 machine
first, etc. Once you have resolved the problem, go back and try to delete the cloud formation stack again.

### Remove secrets

Log in to the AWS console for your account, then go to `Systems Manager > Parameter store`. Here you might find leftover
secrets. 

### Remove deploy key(s) in IAC repository

Log in to Github, open your IAC repository, choose `Settings > Deploy keys`. Delete any unused keys here.

### Clean up infrastructure-as-code(IAC) repository

The following paths reside in the top level directory of your infrastructure-as-code repository:

1. Delete the entry for the cluster in `.okctl.yml`.
1. Delete the directory `infrastructure/<environment name>`.

If the cluster you are deleting is the last one in this IAC repository, you can delete both the `.okctl.yml` file and
the `infrastructure` directory.

### List all AWS resources created by `okctl`

It is possible to list all AWS resources created by `okctl` by running the command below. This can be a useful command to run to find any missing resources.

```bash
aws resourcegroupstaggingapi get-resources \
> --tag-filters Key=alpha.okctl.io/okctl-version,Values=dev \
> --tags-per-page 100
```
