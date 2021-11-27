<?php

namespace App\Jobs;

use Illuminate\Bus\Queueable;
use Illuminate\Contracts\Queue\ShouldBeUnique;
use Illuminate\Contracts\Queue\ShouldQueue;
use Illuminate\Foundation\Bus\Dispatchable;
use Illuminate\Queue\InteractsWithQueue;
use Illuminate\Queue\SerializesModels;
use Illuminate\Support\Facades\Redis;

/**
 * Keep user until removes the user after a
 * defined period of time.
 *
 */
class KeepUserUntil implements ShouldQueue
{
    // Traits
    use Dispatchable, InteractsWithQueue, Queueable, SerializesModels;

    // User name
    protected $name;

    /**
     * Create a new job instance.
     *
     * @return void
     */
    public function __construct($name)
    {
        $this->name = $name;
    }

    /**
     * Execute the job.
     *
     * @return void
     */
    public function handle()
    {
        if (Redis::exists($this->name))
            Redis::del($this->name); // Removing user from online users
    }
}
