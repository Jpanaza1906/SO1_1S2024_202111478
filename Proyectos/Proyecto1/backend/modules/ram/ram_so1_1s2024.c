#include <linux/module.h>
#include <linux/proc_fs.h>
#include <linux/sysinfo.h> // ram 
#include <linux/seq_file.h>
#include <linux/mm.h>

MODULE_LICENSE("GPL");
MODULE_AUTHOR("Jose Panaza");
MODULE_DESCRIPTION("RAM system information");
MODULE_VERSION("1.0");

struct sysinfo inf;

static int write_proc(struct seq_file *file_proc, void *v) {
    unsigned long total_memory, free_memory, used_memory;
    unsigned long percentage;
    si_meminfo(&inf);

    total_memory = inf.totalram * inf.mem_unit;
    used_memory = (inf.totalram - inf.freeram) * inf.mem_unit;
    free_memory = total_memory - used_memory;
    percentage = (used_memory * 100) / total_memory;


    seq_printf(file_proc, "{\"total_memory\": %lu, \"used_memory\": %lu, \"free_memory\": %lu, \"ram_percentage\": %lu}\n", total_memory, used_memory, free_memory, percentage);

    return 0;
}

static int open_proc(struct inode *inode, struct file *file) {
    return single_open(file, write_proc, NULL);
}

// File operations

static struct proc_ops file_ops = {
    .proc_open = open_proc,
    .proc_read = seq_read
};

static int __init start(void) {
    proc_create("ram_so1_1s2024", 0, NULL, &file_ops);
    printk(KERN_INFO "RAM module loaded\n");
    return 0;
}

static void __exit finish(void) {
    remove_proc_entry("ram_so1_1s2024", NULL);
    printk(KERN_INFO "RAM module unloaded\n");
}

module_init(start);
module_exit(finish);

// sudo insmod ram_so1_1s2024.ko
// sudo rmmod ram_so1_1s2024.ko